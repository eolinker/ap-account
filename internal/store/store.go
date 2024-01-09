package store

import (
	"gitlab.eolink.com/apinto/common/autowire"
	"gitlab.eolink.com/apinto/common/store"
	"reflect"
)

type UserAuthStore interface {
	store.IBaseStore[UserAuth]
}
type userAuthStore struct {
	store.BaseStore[UserAuth]
}

type UserInfoStore interface {
	store.IBaseStore[UserInfo]
}
type userInfoStore struct {
	store.BaseStore[UserInfo]
}

type UserLoginLogStore interface {
	store.IBaseStore[UserLoginLog]
}
type userLoginLogStore struct {
	store.BaseStore[UserLoginLog]
}

type UserRegisterLogStore interface {
	store.IBaseStore[UserRegisterLog]
}
type userRegisterLogStore struct {
	store.BaseStore[UserRegisterLog]
}

type UserInfoUpdateLogStore interface {
	store.IBaseStore[UserInfoUpdateLog]
}
type userInfoUpdateLogStore struct {
	store.BaseStore[UserInfoUpdateLog]
}

func init() {
	autowire.Auto[UserAuthStore](func() reflect.Value {
		return reflect.ValueOf(new(userAuthStore))
	})
	autowire.Auto[UserInfoStore](func() reflect.Value {
		return reflect.ValueOf(new(userInfoStore))
	})

	autowire.Auto[UserLoginLogStore](func() reflect.Value {
		return reflect.ValueOf(new(userLoginLogStore))
	})
	autowire.Auto[UserRegisterLogStore](func() reflect.Value {
		return reflect.ValueOf(new(userRegisterLogStore))
	})

	autowire.Auto[UserInfoUpdateLogStore](func() reflect.Value {
		return reflect.ValueOf(new(userInfoUpdateLogStore))
	})
}
