package account

import (
	"context"
	"errors"
)

var (
	ErrorUserNotFound = errors.New("user not found")
)

type UserId = string

type AccountService interface {
	Login(ctx context.Context, driver string, identifier string, certificate string) (UserId, error)
	AddAuth(ctx context.Context, driver string, uid string, identifier string, certificate string) error

	CheckAuth(ctx context.Context, driver string, identifier string, certificate string) (UserId, error)

	Logout(ctx context.Context, driver string, identifier string) error
	GetUserInfo(ctx context.Context, uid UserId) (UserInfo, error)
	UpdateUserInfo(ctx context.Context, uid UserId, userInfo UserInfo, operator UserId) error
	Remove(ctx context.Context, uid UserId) error
}
