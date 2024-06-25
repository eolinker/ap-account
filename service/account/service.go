package account

import (
	"context"
	"reflect"

	"github.com/eolinker/ap-account/service/usage"
	"github.com/eolinker/go-common/autowire"
)

type UserId = string

type IAccountService interface {
	//AddAuth(ctx context.Context, driver string, uid string, identifier string, certificate string) error
	Save(ctx context.Context, driver string, uid string, identifier string, certificate string) error
	//CreateAccount(ctx context.Context, uid string, userInfo account.UserInfo) (UserId, error)

	GetForUser(ctx context.Context, id string) ([]*UserAuth, error)
	GetIdentifier(ctx context.Context, driver string, identifier string) (*UserAuth, error)
	//GetUserInfo(ctx context.Context, uid UserId) (account.UserInfo, error)
	//UpdateUserInfo(ctx context.Context, uid UserId, userInfo account.UserInfo, operator UserId) error
	//Remove(ctx context.Context, uid UserId) error
	usage.IUserUsageService
}

func init() {
	autowire.Auto[IAccountService](func() reflect.Value {
		return reflect.ValueOf(new(imlAccountService))
	})
}
