package member

import (
	"context"
	"reflect"

	user_dto "gitlab.eolink.com/apinto/aoaccount/module/user/dto"
	department_member "gitlab.eolink.com/apinto/aoaccount/service/department-member"
	"gitlab.eolink.com/apinto/aoaccount/service/member"
	"gitlab.eolink.com/apinto/aoaccount/service/user"
	user_group "gitlab.eolink.com/apinto/aoaccount/service/user-group"
	"gitlab.eolink.com/apinto/common/auto"
	"gitlab.eolink.com/apinto/common/autowire"
	"gitlab.eolink.com/apinto/common/utils"
)

var (
	_ IMemberModule = (*imlMemberModule)(nil)
)

type IMemberModule interface {
	UserGroupMember(ctx context.Context, groupId ...string) ([]*user_dto.UserInfo, error)
}

type imlMemberModule struct {
	memberService           user_group.IUserGroupMemberService `autowired:""`
	departmentMemberService department_member.IMemberService   `autowired:""`
	userService             user.IUserService                  `autowired:""`
}

func (m *imlMemberModule) UserGroupMember(ctx context.Context, groupId ...string) ([]*user_dto.UserInfo, error) {

	members, err := m.memberService.Members(ctx, groupId, nil)
	if err != nil {
		return nil, err
	}

	userids := utils.SliceToSlice(members, member.UserID, func(m *member.Member) bool {
		return m.Come != ""
	})

	if len(userids) == 0 {
		return nil, nil
	}
	users, err := m.userService.Get(ctx, userids...)
	if err != nil {
		return nil, err
	}
	result := utils.SliceToSlice(users, user_dto.CreateUserInfoFromModel)

	groups, err := m.memberService.FilterMembersForUser(ctx, userids...)
	if err != nil {
		return nil, err
	}
	departmentMembers, err := m.departmentMemberService.Members(ctx, nil, userids)
	if err != nil {
		return nil, err
	}
	departmentMemberMap := utils.SliceToMapArrayO(utils.SliceToSlice(departmentMembers, func(s *member.Member) *member.Member {
		return s
	}, func(m *member.Member) bool {
		return m.Come != ""
	}), func(t *member.Member) (string, string) {
		return t.UID, t.Come
	})
	for _, r := range result {
		r.Department = auto.List(departmentMemberMap[r.Uid])
		r.UserGroups = auto.List(groups[r.Uid])
	}
	return result, nil
}

func init() {
	autowire.Auto[IMemberModule](func() reflect.Value {
		return reflect.ValueOf(new(imlMemberModule))
	})
}
