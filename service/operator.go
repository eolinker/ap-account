package account

import (
	"context"
	"gitlab.eolink.com/apinto/aoaccount/internal/store"
	"gitlab.eolink.com/apinto/common/autowire"
	"reflect"
)

type UserHandler interface {
	GetUserId() []string
	SetUserName(operators map[string]string)
}
type IUserInfoService interface {
	CompleteUserName(ctx context.Context, handler UserHandler)
}

var (
	_ IUserInfoService = (*imlUserInfoService)(nil)
)

type imlUserInfoService struct {
	store store.IUserInfoStore `autowired:""`
}

func (i *imlUserInfoService) CompleteUserName(ctx context.Context, handler UserHandler) {
	ids := handler.GetUserId()
	if len(ids) == 0 {
		return
	}
	// 查询用户信息
	users, err := i.store.ListQuery(ctx, "id in (?)", []interface{}{ids}, "id asc")
	if err != nil {
		return
	}
	// 填充用户信息
	for _, user := range users {
		handler.SetUserName(map[string]string{
			user.Uid: user.NickName,
		})
	}
}

func init() {
	autowire.Auto[IUserInfoService](func() reflect.Value {
		return reflect.ValueOf(new(imlUserInfoService))
	})
}
