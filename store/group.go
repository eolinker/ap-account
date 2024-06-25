package store

import (
	"github.com/eolinker/ap-account/store/member"
	"github.com/eolinker/go-common/autowire"
	"github.com/eolinker/go-common/store"
	"reflect"
)

type IUserGroupStore interface {
	store.IBaseStore[UserGroup]
}

type IUserGroupMemberStore member.IMemberStore

type imlUserGroupStore struct {
	store.Store[UserGroup]
}

func init() {
	autowire.Auto[IUserGroupStore](func() reflect.Value {
		return reflect.ValueOf(new(imlUserGroupStore))
	})
	autowire.Auto[IUserGroupMemberStore](func() reflect.Value {
		return reflect.ValueOf(member.NewMemberStore("user_group"))
	})
}
