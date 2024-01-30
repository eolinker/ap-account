package store

import (
	"gitlab.eolink.com/apinto/common/autowire"
	"gitlab.eolink.com/apinto/common/store"
	"reflect"
)

var (
	_ IUserAuthStore        = (*userAuthStore)(nil)
	_ IUserInfoStore        = (*userInfoStore)(nil)
	_ IUserLoginLogStore    = (*userLoginLogStore)(nil)
	_ IUserRegisterLogStore = (*userRegisterLogStore)(nil)
	_ IUserGroupStore       = (*imlUserGroupStore)(nil)
)

type IUserAuthStore interface {
	store.IBaseStore[UserAuth]
}
type userAuthStore struct {
	store.Store[UserAuth]
}

type IUserInfoStore interface {
	store.IBaseStore[UserInfo]
	//SearchWithDepartment(ctx context.Context, department, keyword string) ([]*UserInfo, error)
}
type userInfoStore struct {
	store.Store[UserInfo]
}

type IUserLoginLogStore interface {
	store.IBaseStore[UserLoginLog]
}
type userLoginLogStore struct {
	store.Store[UserLoginLog]
}

type IUserRegisterLogStore interface {
	store.IBaseStore[UserRegisterLog]
}
type userRegisterLogStore struct {
	store.Store[UserRegisterLog]
}

type IUserInfoUpdateLogStore interface {
	store.IBaseStore[UserInfoUpdateLog]
}
type userInfoUpdateLogStore struct {
	store.Store[UserInfoUpdateLog]
}

func init() {
	autowire.Auto[IUserAuthStore](func() reflect.Value {
		return reflect.ValueOf(new(userAuthStore))
	})
	autowire.Auto[IUserInfoStore](func() reflect.Value {
		return reflect.ValueOf(new(userInfoStore))
	})

	autowire.Auto[IUserLoginLogStore](func() reflect.Value {
		return reflect.ValueOf(new(userLoginLogStore))
	})
	autowire.Auto[IUserRegisterLogStore](func() reflect.Value {
		return reflect.ValueOf(new(userRegisterLogStore))
	})

	autowire.Auto[IUserInfoUpdateLogStore](func() reflect.Value {
		return reflect.ValueOf(new(userInfoUpdateLogStore))
	})
}
