package user

import (
	"context"
	auth_password "gitlab.eolink.com/apinto/aoaccount/auth_driver/auth-password"
	user_dto "gitlab.eolink.com/apinto/aoaccount/module/user/dto"
	department_member "gitlab.eolink.com/apinto/aoaccount/service/department-member"
	"gitlab.eolink.com/apinto/aoaccount/service/user"
	"gitlab.eolink.com/apinto/common/store"
	"gitlab.eolink.com/apinto/common/utils"
)

var (
	_ IUserModule = (*imlUserModule)(nil)
)

const (
	defaultInitPassword = "12345678"
)

type imlUserModule struct {
	userService   user.IUserService                `autowired:""`
	memberService department_member.IMemberService `autowired:""`
	authPassword  auth_password.AuthPassword       `autowired:""`
	transaction   store.ITransaction               `autowired:""`
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
	list, err := s.userService.Search(ctx, department, keyword)
	if err != nil {
		return nil, err
	}
	result := utils.SliceToSlice(list, user_dto.CreateUserInfoFromModel)
	return result, nil
}

func (s *imlUserModule) AddForPassword(ctx context.Context, user *user_dto.CreateUser) (string, error) {
	newId := ""
	err := s.transaction.Transaction(ctx, func(txCtx context.Context) error {
		newUser, err := s.userService.Create(ctx, "", user.Name, user.Email, user.Mobile)
		if err != nil {
			return err
		}

		err = s.authPassword.Save(ctx, newUser.UID, user.Name, defaultInitPassword)
		if err != nil {
			return err
		}

		if len(user.Departments) > 0 {
			for _, department := range user.Departments {
				_ = s.memberService.AddMembers(ctx, department, newUser.UID)
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
	return s.userService.Delete(ctx, ids...)

}
