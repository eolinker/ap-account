package user_group

import (
	"context"
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
}

type IUserGroupMemberService interface {
	AddGroup(ctx context.Context, groupID string, userids ...string) error
	RemoveGroup(ctx context.Context, groupID, userID string) error
	Members(ctx context.Context, gids ...string) ([]*Member, error)
}

func init() {
	autowire.Auto[IUserGroupService](func() reflect.Value {
		return reflect.ValueOf(new(imlUserGroupService))
	})
	autowire.Auto[IUserGroupMemberService](func() reflect.Value {
		return reflect.ValueOf(new(imlUserMemberService))
	})
}
