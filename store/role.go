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

func init() {
	autowire.Auto[IRoleStore](func() reflect.Value {
		return reflect.ValueOf(new(imlRoleStore))
	})
}
