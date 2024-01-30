package account

import (
	"context"
	"errors"
	"gitlab.eolink.com/apinto/aoaccount/service/usage"
	"gitlab.eolink.com/apinto/common/autowire"
	"reflect"
)

var (
	ErrorUserNotFound = errors.New("user not found")
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
