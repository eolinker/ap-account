package department

import (
	"context"
	"errors"
	"github.com/google/uuid"
	department_dto "gitlab.eolink.com/apinto/aoaccount/module/department/dto"
	"gitlab.eolink.com/apinto/aoaccount/service/account"
	"gitlab.eolink.com/apinto/aoaccount/service/department"
	department_member "gitlab.eolink.com/apinto/aoaccount/service/department-member"
	"gitlab.eolink.com/apinto/common/utils"
)

var (
	_ IDepartmentModule = (*imlDepartmentModule)(nil)
)

type imlDepartmentModule struct {
	service       department.IDepartmentService    `autowired:""`
	userService   account.IAccountService          `autowired:""`
	memberService department_member.IMemberService `autowired:""`
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

func (m *imlDepartmentModule) Tree(ctx context.Context) (*department_dto.Department, int, error) {
	list, err := m.service.Get(ctx)
	if err != nil {
		return nil, 0, err
	}
	members, err := m.memberService.Members(ctx)
	if err != nil {
		return nil, 0, err
	}
	departmentsMembers := utils.SliceToMapArrayO(members, func(t *department_member.Member) (string, string) {
		return t.Department, t.User
	})
	nodes := utils.SliceToMapO(list, func(s *department.Department) (string, *Node) {
		return s.Id, &Node{
			Id:       s.Id,
			Name:     s.Name,
			ParentId: s.ParentId,
			Children: nil,
			Members:  utils.NewSet(departmentsMembers[s.Id]...),
		}
	})
	root := &Node{
		Id:       "Root",
		Name:     "Root",
		ParentId: "",
		Members:  utils.NewSet[string](),
		Children: nil,
	}
	for _, n := range nodes {
		if n.ParentId != "" {
			if p, has := nodes[n.ParentId]; has {
				p.Children = append(p.Children, n)
				continue
			}
		}
		root.Children = append(root.Children, n)
	}
	root.SMembers()
	unKnownDepartments := len(departmentsMembers[""])
	return root.toDto(), unKnownDepartments, nil

}

func (m *imlDepartmentModule) AddMember(ctx context.Context, id string, member *department_dto.AddMember) error {
	return m.memberService.AddMembers(ctx, id, member.UserIds...)
}

func (m *imlDepartmentModule) RemoveMember(ctx context.Context, id string, uid string) error {
	return m.memberService.RemoveMembers(ctx, id, uid)
}

func (m *imlDepartmentModule) RemoveMembers(ctx context.Context, id string, members *department_dto.RemoveMember) error {
	return m.memberService.RemoveMembers(ctx, id, members.UserIds...)
}
