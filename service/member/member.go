package member

import (
	"context"

	"gitlab.eolink.com/apinto/aoaccount/service/usage"
	"gitlab.eolink.com/apinto/aoaccount/store/member"
	"gitlab.eolink.com/apinto/common/utils"
)

var (
	_ IMemberService = (*Service[member.IMemberStore])(nil)
)

type IMemberService interface {
	AddMemberTo(ctx context.Context, cid string, userIds ...string) error
	RemoveMemberFrom(ctx context.Context, cid string, uuids ...string) error
	Members(ctx context.Context, cids []string, users []string) ([]*Member, error)
	FilterMembersForUser(ctx context.Context, users ...string) (map[string][]string, error)
	Delete(ctx context.Context, cid ...string) error
	usage.IUserUsageService
}

type Service[S member.IMemberStore] struct {
	store S `autowired:""`
}

func (s *Service[S]) FilterMembersForUser(ctx context.Context, users ...string) (map[string][]string, error) {
	members, err := s.Members(ctx, nil, users)
	if err != nil {
		return nil, err
	}

	membersMap := make(map[string][]string)
	for _, v := range members {
		if v.Come == "" {
			continue
		}
		membersMap[v.UID] = append(membersMap[v.UID], v.Come)
	}
	return membersMap, nil
}

func (s *Service[S]) OnComplete() {
	usage.RegisterUser(s)
}

func (s *Service[S]) AddMemberTo(ctx context.Context, cid string, userIds ...string) error {
	return s.store.AddMember(ctx, cid, userIds...)
}

func (s *Service[S]) RemoveMemberFrom(ctx context.Context, cid string, userIds ...string) error {
	return s.store.RemoveMember(ctx, cid, userIds...)
}

func (s *Service[S]) Members(ctx context.Context, cids []string, users []string) ([]*Member, error) {

	members, err := s.store.Members(ctx, cids, users)
	if err != nil {
		return nil, err
	}
	return utils.SliceToSlice(members, toModel), nil
}

func (s *Service[S]) Delete(ctx context.Context, cid ...string) error {
	return s.store.Delete(ctx, cid...)
}

func (s *Service[S]) OnRemoveUsers(ctx context.Context, ids ...string) error {
	return s.store.RemoveUser(ctx, ids...)
}
