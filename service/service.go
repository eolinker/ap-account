package account

import (
	"context"
	"errors"
	"gitlab.eolink.com/apinto/common/autowire"
	"reflect"
)

var (
	ErrorUserNotFound = errors.New("user not found")
)

type UserId = string

type IAccountService interface {
	AddAuth(ctx context.Context, driver string, uid string, identifier string, certificate string) error
	CreateAccount(ctx context.Context, uid string, userInfo UserInfo) (UserId, error)
	CheckAuth(ctx context.Context, driver string, identifier string, certificate string) (UserId, error)

	GetUserInfo(ctx context.Context, uid UserId) (UserInfo, error)
	UpdateUserInfo(ctx context.Context, uid UserId, userInfo UserInfo, operator UserId) error
	Remove(ctx context.Context, uid UserId) error
}

func init() {
	autowire.Auto[IAccountService](func() reflect.Value {
		return reflect.ValueOf(new(imlAccountService))
	})
}
