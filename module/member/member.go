package member

import (
	"context"
	user_dto "gitlab.eolink.com/apinto/aoaccount/module/user/dto"
	"gitlab.eolink.com/apinto/aoaccount/service/user"
	user_group "gitlab.eolink.com/apinto/aoaccount/service/user-group"
	"gitlab.eolink.com/apinto/common/auto"
	"gitlab.eolink.com/apinto/common/autowire"
	"gitlab.eolink.com/apinto/common/utils"
	"reflect"
)

var (
	_ IMemberModule = (*imlMemberModule)(nil)
)

type IMemberModule interface {
	UserGroupMember(ctx context.Context, groupId ...string) ([]*user_dto.UserInfo, error)
}

type imlMemberModule struct {
	memberService user_group.IUserGroupMemberService `autowired:""`
	userService   user.IUserService                  `autowired:""`
}

func (m *imlMemberModule) UserGroupMember(ctx context.Context, groupId ...string) ([]*user_dto.UserInfo, error) {

	members, err := m.memberService.Members(ctx, groupId...)
	if err != nil {
		return nil, err
	}
	userids := utils.SliceToSlice(members, func(s *user_group.Member) string {
		return s.UserId
	})

	users, err := m.userService.Get(ctx, userids...)
	if err != nil {
		return nil, err
	}
	result := utils.SliceToSlice(users, user_dto.CreateUserInfoFromModel)
	auto.CompleteLabels(ctx, result)
	return result, nil
}

func init() {
	autowire.Auto[IMemberModule](func() reflect.Value {
		return reflect.ValueOf(new(imlMemberModule))
	})
}
