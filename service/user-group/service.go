package user_group

import (
	"context"
	"gitlab.eolink.com/apinto/aoaccount/service/member"
	"gitlab.eolink.com/apinto/aoaccount/store"
	"gitlab.eolink.com/apinto/common/autowire"
	"reflect"
)

type IUserGroupService interface {
	Crete(ctx context.Context, id, name string) error
	Edit(ctx context.Context, id, name string) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*UserGroup, error)
	// GetList description of the Go function.
	//
	// ctx context.Context
	// []*UserGroup, error
	GetList(ctx context.Context) ([]*UserGroup, error)
	Search(ctx context.Context, keyword string) ([]*UserGroup, error)
}

type IUserGroupMemberService member.IMemberService

func init() {
	autowire.Auto[IUserGroupService](func() reflect.Value {
		return reflect.ValueOf(new(imlUserGroupService))
	})
	autowire.Auto[IUserGroupMemberService](func() reflect.Value {
		return reflect.ValueOf(new(member.Service[store.IUserGroupMemberStore]))
	})
}
