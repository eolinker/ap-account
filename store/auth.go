package store

import (
	"reflect"

	"github.com/eolinker/go-common/autowire"
	"github.com/eolinker/go-common/store"
)

type IAuthDriverStore interface {
	store.IBaseStore[AuthDriver]
}
type authDriverStore struct {
	store.Store[AuthDriver]
}

func init() {
	autowire.Auto[IAuthDriverStore](func() reflect.Value {
		return reflect.ValueOf(new(authDriverStore))
	})
}
