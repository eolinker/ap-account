package department

import (
	"context"
	"errors"

	"github.com/eolinker/go-common/store"

	department_dto "github.com/eolinker/ap-account/module/department/dto"
	"github.com/eolinker/ap-account/service/account"
	"github.com/eolinker/ap-account/service/department"
	department_member "github.com/eolinker/ap-account/service/department-member"
	"github.com/eolinker/go-common/utils"
	"github.com/google/uuid"
)

var (
	_ IDepartmentModule = (*imlDepartmentModule)(nil)
)

type imlDepartmentModule struct {
	service       department.IDepartmentService    `autowired:""`
	userService   account.IAccountService          `autowired:""`
	memberService department_member.IMemberService `autowired:""`
	transaction   store.ITransaction               `autowired:""`
}

func (m *imlDepartmentModule) CreateDepartment(ctx context.Context, department *department_dto.Create) (string, error) {
	id := department.Id
	if id == "" {
		id = uuid.NewString()
	}
	err := m.service.Create(ctx, id, department.Name, department.ParentID)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (m *imlDepartmentModule) EditDepartment(ctx context.Context, id string, department *department_dto.Edit) error {
	if id == "" && (department == nil || department.Id == nil || *(department.Id) == "") {
		return errors.New("id is required")
	}
	if id == "" {
		id = *department.Id
	}
	if id == "" {
		return errors.New("id is required")
	}
	return m.service.Edit(ctx, id, department.Name, department.ParentId)

}

func (m *imlDepartmentModule) Delete(ctx context.Context, id string) error {
	return m.service.Delete(ctx, id)
}

func (m *imlDepartmentModule) Simple(ctx context.Context) (*department_dto.Simple, error) {
	list, err := m.service.Get(ctx)
	if err != nil {
		return nil, err
	}
	nodes := utils.SliceToMapO(list, func(s *department.Department) (string, *department_dto.Simple) {
		return s.Id, &department_dto.Simple{
			Id:     s.Id,
			Name:   s.Name,
			Parent: s.ParentId,
		}
	})
	root := &department_dto.Simple{
		Id:   "",
		Name: "",
	}
	for _, s := range nodes {
		if p, has := nodes[s.Parent]; has {
			p.Children = append(p.Children, s)
			continue
		}
		root.Children = append(root.Children, s)
	}
	return root, nil
}

func (m *imlDepartmentModule) Tree(ctx context.Context) (*department_dto.Department, error) {
	tree, err := m.service.Tree(ctx)
	if err != nil {
		return nil, err
	}
	members, err := m.memberService.Members(ctx, nil, nil)
	if err != nil {
		return nil, err
	}

	departmentsMembers := utils.SliceToMapArrayO(members, func(t *department_member.Member) (string, string) {
		return t.Come, t.UID
	})

	return toDto(tree, departmentsMembers), nil

}

func (m *imlDepartmentModule) AddMember(ctx context.Context, member *department_dto.AddMember) error {
	return m.transaction.Transaction(ctx, func(txCtx context.Context) error {
		for _, cid := range member.DepartmentIds {
			err := m.memberService.AddMemberTo(ctx, cid, member.UserIds...)
			if err != nil {
				return err
			}
		}
		return nil
	})

}

func (m *imlDepartmentModule) RemoveMember(ctx context.Context, id string, uid string) error {
	return m.memberService.RemoveMemberFrom(ctx, id, uid)
}

func (m *imlDepartmentModule) RemoveMembers(ctx context.Context, id string, members *department_dto.RemoveMember) error {
	return m.memberService.RemoveMemberFrom(ctx, id, members.UserIds...)
}
