package store

import (
	"gitlab.eolink.com/apinto/aoaccount/store/member"
	"gitlab.eolink.com/apinto/common/autowire"
	"gitlab.eolink.com/apinto/common/store"
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
