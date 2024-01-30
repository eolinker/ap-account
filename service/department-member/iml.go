package department_member

import (
	"context"
	"errors"
	"gitlab.eolink.com/apinto/aoaccount/service/usage"
	"gitlab.eolink.com/apinto/aoaccount/store"
	"gitlab.eolink.com/apinto/common/autowire"
	"gitlab.eolink.com/apinto/common/utils"
	"gorm.io/gorm"
	"time"
)

var (
	_ IMemberService          = (*imlMemberService)(nil)
	_ autowire.Complete       = (*imlMemberService)(nil)
	_ usage.IUserUsageService = (*imlMemberService)(nil)
)

type imlMemberService struct {
	memberStore     store.DepartmentMemberStore `autowired:""`
	departmentStore store.DepartmentStore       `autowired:""`
}

func (s *imlMemberService) OnComplete() {
	usage.RegisterUser(s)
}

func (s *imlMemberService) RemoveUser(ctx context.Context, ids ...string) error {
	_, err := s.memberStore.DeleteQuery(ctx, "uid in (?)", ids)
	if err != nil {
		return err
	}
	return nil
}

func (s *imlMemberService) AddMembers(ctx context.Context, departmentId string, userIds ...string) error {
	if len(userIds) == 0 {
		return errors.New("invalid user")
	}

	mbs := make([]*store.DepartmentMember, 0, len(userIds))
	for _, userId := range userIds {
		mbs = append(mbs, &store.DepartmentMember{
			Id:         0,
			Department: departmentId,
			Uid:        userId,
			CreateTime: time.Time{},
		})
	}
	return s.memberStore.Transaction(ctx, func(txCtx context.Context) error {
		dv, err := s.departmentStore.First(ctx, map[string]interface{}{
			"uuid": departmentId,
		})
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if dv == nil {
			return errors.New("department not exist")
		}
		for _, m := range mbs {
			err := s.memberStore.Save(ctx, m)
			if err != nil {
				return err
			}
		}
		return nil
	})

}

func (s *imlMemberService) RemoveMembers(ctx context.Context, departmentId string, userIds ...string) error {
	query, err := s.memberStore.DeleteQuery(ctx, "`department` = ? AND  `uid` in(?)", []interface{}{departmentId, userIds})
	if err != nil {
		return err
	}
	if query == 0 {
		return errors.New("no members")
	}
	return nil
}

func (s *imlMemberService) Members(ctx context.Context) ([]*Member, error) {
	members, err := s.memberStore.UserMembers(ctx)
	if err != nil {
		return nil, err
	}
	return utils.SliceToSlice(members, func(t *store.UserMember) *Member {
		return &Member{
			User:       t.Uid,
			Department: t.Department,
		}
	}), nil
}
