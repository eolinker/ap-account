package store

import (
	"github.com/eolinker/ap-account/store/member"
	"github.com/eolinker/go-common/autowire"
	"github.com/eolinker/go-common/store"
	"reflect"
)

var (
	_ DepartmentStore = (*imlDepartmentStore)(nil)
)

type DepartmentStore interface {
	store.IBaseStore[Department]
}
type DepartmentMemberStore member.IMemberStore

type imlDepartmentStore struct {
	store.Store[Department]
}

func init() {
	autowire.Auto[DepartmentStore](func() reflect.Value {
		return reflect.ValueOf(new(imlDepartmentStore))
	})
	autowire.Auto[DepartmentMemberStore](func() reflect.Value {
		return reflect.ValueOf(member.NewMemberStore("department"))
	})
}
