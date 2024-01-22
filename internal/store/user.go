package store

import (
	"gitlab.eolink.com/apinto/common/autowire"
	"gitlab.eolink.com/apinto/common/store"
	"reflect"
)

type IUserAuthStore interface {
	store.IBaseStore[UserAuth]
}
type userAuthStore struct {
	store.BaseStore[UserAuth]
}

type IUserInfoStore interface {
	store.IBaseStore[UserInfo]
}
type userInfoStore struct {
	store.BaseStore[UserInfo]
}

type IUserLoginLogStore interface {
	store.IBaseStore[UserLoginLog]
}
type userLoginLogStore struct {
	store.BaseStore[UserLoginLog]
}

type IUserRegisterLogStore interface {
	store.IBaseStore[UserRegisterLog]
}
type userRegisterLogStore struct {
	store.BaseStore[UserRegisterLog]
}

type IUserInfoUpdateLogStore interface {
	store.IBaseStore[UserInfoUpdateLog]
}
type userInfoUpdateLogStore struct {
	store.BaseStore[UserInfoUpdateLog]
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
