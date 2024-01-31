package store

import (
	"gitlab.eolink.com/apinto/aoaccount/store/member"
	"gitlab.eolink.com/apinto/common/autowire"
	"gitlab.eolink.com/apinto/common/store"
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
