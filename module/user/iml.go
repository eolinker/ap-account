package user

import (
	"context"
	"fmt"
	"github.com/eolinker/ap-account/service/department"
	"log"

	auth_password "github.com/eolinker/ap-account/auth_driver/auth-password"
	user_dto "github.com/eolinker/ap-account/module/user/dto"
	department_member "github.com/eolinker/ap-account/service/department-member"
	"github.com/eolinker/ap-account/service/user"
	user_group "github.com/eolinker/ap-account/service/user-group"
	"github.com/eolinker/go-common/auto"
	"github.com/eolinker/go-common/register"
	"github.com/eolinker/go-common/server"
	"github.com/eolinker/go-common/store"
	"github.com/eolinker/go-common/utils"
)

var (
	_ IUserModule = (*imlUserModule)(nil)
)

const (
	defaultInitPassword = "12345678"
)

type imlUserModule struct {
	userService             user.IUserService                  `autowired:""`
	departmentMemberService department_member.IMemberService   `autowired:""`
	departmentService       department.IDepartmentService      `autowired:""`
	authPassword            auth_password.AuthPassword         `autowired:""`
	userGroupsMemberService user_group.IUserGroupMemberService `autowired:""`
	transaction             store.ITransaction                 `autowired:""`
}

func (s *imlUserModule) UpdateInfo(ctx context.Context, id string, user *user_dto.EditUser) error {
	if user == nil {
		return nil
	}
	_, err := s.userService.Update(ctx, id, user.Name, nil, nil)
	return err
}

func (s *imlUserModule) Simple(ctx context.Context, keyword string) ([]*user_dto.UserSimple, error) {
	list, err := s.userService.Search(ctx, keyword, -1)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	result := utils.SliceToSlice(list, func(s *user.User) *user_dto.UserSimple {
		return &user_dto.UserSimple{
			Uid:        s.UID,
			Name:       s.Username,
			Email:      s.Email,
			Department: nil,
			UserGroups: nil,
		}
	}, func(u *user.User) bool {
		return u.Status == 1
	})
	userIds := utils.SliceToSlice(list, func(m *user.User) string {
		return m.UID
	})
	members, err := s.departmentMemberService.FilterMembersForUser(ctx, userIds...)
	if err != nil {
		return nil, err
	}
	groups, err := s.userGroupsMemberService.FilterMembersForUser(ctx, userIds...)
	if err != nil {
		return nil, err
	}

	for _, r := range result {
		r.Department = auto.List(members[r.Uid])
		r.UserGroups = auto.List(groups[r.Uid])
	}
	return result, nil
}

func (s *imlUserModule) OnComplete() {
	register.Handle(func(v server.Server) {
		ctx := context.Background()
		users, err := s.userService.Get(ctx, "admin")
		if err != nil {
			return
		}
		if len(users) == 0 {
			err := s.transaction.Transaction(ctx, func(ctx context.Context) error {
				create, err := s.userService.Create(ctx, "admin", "admin", "", "")
				if err != nil {
					return err
				}
				return s.authPassword.Save(ctx, create.UID, "admin", defaultInitPassword)

			})
			if err != nil {
				log.Fatal("init admin error: ", err.Error())
			}
		}
	})
}

func (s *imlUserModule) CountStatus(ctx context.Context, enable bool) (int, error) {
	status := 0
	if enable {
		status = 1
	} else {
		status = 2
	}
	count, err := s.userService.CountStatus(ctx, status)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (s *imlUserModule) Search(ctx context.Context, department string, keyword string) ([]*user_dto.UserInfo, error) {

	var list []*user.User
	var err error

	switch department {
	case "disable":
		list, err = s.userService.Search(ctx, keyword, 1)
	case "unknown":
		list, err = s.userService.SearchUnknown(ctx, keyword)
	case "":
		list, err = s.userService.Search(ctx, keyword, -1)
	default:
		tree, errT := s.departmentService.Tree(ctx)
		if errT != nil {
			return nil, errT
		}
		if node, has := tree.Find(department); has {
			list, err = s.userService.Search(ctx, keyword, -1, node.GetChildren()...)
		} else {
			return nil, fmt.Errorf("departemnt %s not exist", department)
		}
	}
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	result := utils.SliceToSlice(list, user_dto.CreateUserInfoFromModel)
	userIds := utils.SliceToSlice(list, func(m *user.User) string {
		return m.UID
	})
	members, err := s.departmentMemberService.FilterMembersForUser(ctx, userIds...)
	if err != nil {
		return nil, err
	}
	groups, err := s.userGroupsMemberService.FilterMembersForUser(ctx, userIds...)
	if err != nil {
		return nil, err
	}

	for _, r := range result {
		r.Department = auto.List(members[r.Uid])
		r.UserGroups = auto.List(groups[r.Uid])
	}
	return result, nil
}

func (s *imlUserModule) AddForPassword(ctx context.Context, user *user_dto.CreateUser) (string, error) {
	newId := ""

	err := s.transaction.Transaction(ctx, func(ctx context.Context) error {
		newUser, err := s.userService.Create(ctx, "", user.Name, user.Email, user.Mobile)
		if err != nil {
			return err
		}

		err = s.authPassword.Save(ctx, newUser.UID, user.Name, defaultInitPassword)
		if err != nil {
			return err
		}

		if len(user.Departments) > 0 {
			for _, dp := range user.Departments {
				err := s.departmentMemberService.AddMemberTo(ctx, dp, newUser.UID)
				if err != nil {
					return err
				}
			}
		}
		newId = newUser.UID
		return nil
	})
	if err != nil {
		return "", err
	}
	return newId, nil
}

func (s *imlUserModule) Disable(ctx context.Context, user *user_dto.Disable) error {
	return s.userService.SetStatus(ctx, 2, user.Users...)
}

func (s *imlUserModule) Enable(ctx context.Context, user *user_dto.Enable) error {
	return s.userService.SetStatus(ctx, 1, user.Users...)

}
func (s *imlUserModule) Delete(ctx context.Context, ids ...string) error {
	return s.transaction.Transaction(ctx, func(txCtx context.Context) error {
		err := s.departmentMemberService.OnRemoveUsers(ctx, ids...)
		if err != nil {
			return err
		}
		err = s.authPassword.Delete(ctx, ids...)
		if err != nil {
			return err
		}
		return s.userService.Delete(ctx, ids...)
	})

}
