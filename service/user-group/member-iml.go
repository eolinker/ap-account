package user_group

import (
	"context"
	"gitlab.eolink.com/apinto/aoaccount/service/usage"
	"gitlab.eolink.com/apinto/aoaccount/store"
	"gitlab.eolink.com/apinto/common/autowire"
	"gitlab.eolink.com/apinto/common/utils"
	"time"
)

var (
	_ IUserGroupMemberService = (*imlUserMemberService)(nil)
	_ usage.IUserUsageService = (*imlUserMemberService)(nil)
	_ autowire.Complete       = (*imlUserMemberService)(nil)
)

type imlUserMemberService struct {
	store store.IUserGroupMemberStore `autowired:""`
}

func (s *imlUserMemberService) OnComplete() {
	usage.RegisterUser(s)
}

func (s *imlUserMemberService) RemoveUser(ctx context.Context, ids ...string) error {
	_, err := s.store.DeleteQuery(ctx, "uid in (?)", []interface{}{ids})
	if err != nil {
		return err
	}
	return nil
}

func (s *imlUserMemberService) AddGroup(ctx context.Context, groupID string, userIds ...string) error {

	for _, uid := range userIds {
		err := s.store.Save(ctx, &store.UserGroupMember{
			Id:         0,
			Gid:        groupID,
			Uid:        uid,
			CreateTime: time.Now(),
		})
		return err
	}
	return nil
}

func (s *imlUserMemberService) RemoveGroup(ctx context.Context, userID, groupID string) error {
	_, err := s.store.DeleteQuery(ctx, "gid = ? and uid = ?", []interface{}{userID, groupID})
	return err
}

func (s *imlUserMemberService) Members(ctx context.Context, gids ...string) ([]*Member, error) {
	if len(gids) == 0 {
		list, err := s.store.List(ctx, map[string]interface{}{})
		if err != nil {
			return nil, err
		}
		return utils.SliceToSlice(list, func(member *store.UserGroupMember) *Member {
			return &Member{
				GroupId: member.Gid,
				UserId:  member.Uid,
			}
		}), nil
	}

	list, err := s.store.ListQuery(ctx, "gid in (?)", []interface{}{gids}, "created_time desc")
	if err != nil {
		return nil, err
	}
	return utils.SliceToSlice(list, func(member *store.UserGroupMember) *Member {
		return &Member{
			UserId:  member.Uid,
			GroupId: member.Gid,
		}
	}), nil

}
