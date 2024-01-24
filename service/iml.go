package account

import (
	"context"
	"gitlab.eolink.com/apinto/aoaccount/internal/store"
	"gitlab.eolink.com/apinto/common/auto"
	"gitlab.eolink.com/apinto/common/utils"
)

var (
	_ IAccountService = (*imlAccountService)(nil)
)

type imlAccountService struct {
	store store.IUserInfoStore `autowired:""`
}

func (s *imlAccountService) OnComplete() {
	auto.RegisterService("user", s)
}

func (s *imlAccountService) Login(ctx context.Context, driver string, identifier string, certificate string) (UserId, error) {
	//TODO implement me
	panic("implement me")
}

func (s *imlAccountService) AddAuth(ctx context.Context, driver string, uid string, identifier string, certificate string) error {
	//TODO implement me
	panic("implement me")
}

func (s *imlAccountService) CheckAuth(ctx context.Context, driver string, identifier string, certificate string) (UserId, error) {
	//TODO implement me
	panic("implement me")
}

func (s *imlAccountService) Logout(ctx context.Context, driver string, identifier string) error {
	//TODO implement me
	panic("implement me")
}

func (s *imlAccountService) GetUserInfo(ctx context.Context, uid UserId) (UserInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (s *imlAccountService) UpdateUserInfo(ctx context.Context, uid UserId, userInfo UserInfo, operator UserId) error {
	//TODO implement me
	panic("implement me")
}

func (s *imlAccountService) Remove(ctx context.Context, uid UserId) error {
	//TODO implement me
	panic("implement me")
}

func (s *imlAccountService) GetLabels(ctx context.Context, ids ...string) map[string]string {
	users, err := s.store.ListQuery(ctx, "uid in (?)", []interface{}{ids}, "id asc")
	if err != nil {
		return make(map[string]string)
	}
	return utils.SliceToMapO(users, func(user *store.UserInfo) (string, string) {
		return user.Uid, user.NickName
	})
}
