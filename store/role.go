package store

import (
	"reflect"

	"github.com/eolinker/go-common/autowire"
	"github.com/eolinker/go-common/store"
)

type IRoleStore interface {
	store.ISearchStore[Role]
}

type imlRoleStore struct {
	store.SearchStore[Role]
}

type IRoleMemberStore interface {
	store.IBaseStore[RoleMember]
}

type imlRoleMemberStore struct {
	store.Store[RoleMember]
}

func init() {
	autowire.Auto[IRoleStore](func() reflect.Value {
		return reflect.ValueOf(new(imlRoleStore))
	})

	autowire.Auto[IRoleMemberStore](func() reflect.Value {
		return reflect.ValueOf(new(imlRoleMemberStore))
	})
}
