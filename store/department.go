package store

import (
	"context"
	"gitlab.eolink.com/apinto/common/autowire"
	"gitlab.eolink.com/apinto/common/store"
	"reflect"
)

var (
	_ DepartmentStore       = (*imlDepartmentStore)(nil)
	_ DepartmentMemberStore = (*imlDepartmentMemberStore)(nil)
)

type DepartmentStore interface {
	store.IBaseStore[Department]
}
type DepartmentMemberStore interface {
	store.IBaseStore[DepartmentMember]
	UserMembers(ctx context.Context) ([]*UserMember, error)
}

type imlDepartmentStore struct {
	store.Store[Department]
}
type imlDepartmentMemberStore struct {
	store.Store[DepartmentMember]
}

func (s *imlDepartmentMemberStore) UserMembers(ctx context.Context) ([]*UserMember, error) {
	result := make([]*UserMember, 0)
	err := s.DB(ctx).Table("user_info").Select("user_info.uid as uid,member.department as department").
		Joins("left join department_member as member on member.uid = user_info.uid and user_info.is_delete = 0").Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func init() {
	autowire.Auto[DepartmentStore](func() reflect.Value {
		return reflect.ValueOf(new(imlDepartmentStore))
	})
	autowire.Auto[DepartmentMemberStore](func() reflect.Value {
		return reflect.ValueOf(new(imlDepartmentMemberStore))
	})
}
