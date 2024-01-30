package user_group

import (
	"context"
	user_group_dto "gitlab.eolink.com/apinto/aoaccount/module/user-group/dto"
	user_group "gitlab.eolink.com/apinto/aoaccount/service/user-group"
	"gitlab.eolink.com/apinto/common/auto"
	"gitlab.eolink.com/apinto/common/utils"
)

var (
	_ IUserGroupModule = (*imlUserGroupModule)(nil)
)

type imlUserGroupModule struct {
	service       user_group.IUserGroupService       `autowired:""`
	memberService user_group.IUserGroupMemberService `autowired:""`
}

func (m *imlUserGroupModule) AddMember(ctx context.Context, id string, member *user_group_dto.AddMember) error {
	return m.memberService.AddGroup(ctx, id, member.Users...)
}

func (m *imlUserGroupModule) RemoveMember(ctx context.Context, id string, uid string) error {
	return m.memberService.RemoveGroup(ctx, id, uid)
}

func (m *imlUserGroupModule) Get(ctx context.Context, id string) (*user_group_dto.UserGroup, error) {
	v, err := m.service.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user_group_dto.UserGroup{
		Id:         v.Id,
		Name:       v.Name,
		Usage:      0,
		Creator:    auto.UUID(v.Creator),
		CreateTime: auto.TimeLabel(v.CreateTime),
	}, nil
}

func (m *imlUserGroupModule) List(ctx context.Context) ([]*user_group_dto.UserGroup, error) {
	list, err := m.service.GetList(ctx)
	if err != nil {
		return nil, err
	}

	members, err := m.memberService.Members(ctx)
	if err != nil {
		return nil, err
	}
	mbsMap := utils.SliceToMapArray(members, func(t *user_group.Member) string {
		return t.GroupId
	})
	result := utils.SliceToSlice(list, func(s *user_group.UserGroup) *user_group_dto.UserGroup {
		return &user_group_dto.UserGroup{
			Id:         s.Id,
			Name:       s.Name,
			Usage:      len(mbsMap[s.Id]),
			Creator:    auto.UUID(s.Creator),
			CreateTime: auto.TimeLabel(s.CreateTime),
		}
	})
	return result, nil
}

// Create description of the Go function.
//
// ctx context.Context, id string, input *user_group_dto.Create
// error
func (m *imlUserGroupModule) Create(ctx context.Context, id string, input *user_group_dto.Create) error {
	return m.service.Crete(ctx, id, input.Name)
}

func (m *imlUserGroupModule) Edit(ctx context.Context, id string, input *user_group_dto.Edit) error {
	return m.service.Edit(ctx, id, input.Name)
}

// Delete description of the Go function.
//
// ctx context.Context, id string.
// error.
func (m *imlUserGroupModule) Delete(ctx context.Context, id string) error {
	return m.service.Delete(ctx, id)
}

// Simple retrieves a list of simple user group DTOs.
//
// ctx context.Context
// []*user_group_dto.Simple, error
func (m *imlUserGroupModule) Simple(ctx context.Context) ([]*user_group_dto.Simple, error) {
	list, err := m.service.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return utils.SliceToSlice(list, func(s *user_group.UserGroup) *user_group_dto.Simple {
		return &user_group_dto.Simple{
			Id:   s.Id,
			Name: s.Name,
		}
	}), nil
}
