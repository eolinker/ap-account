package user_group

import (
	"context"
	user_group_dto "github.com/eolinker/ap-account/module/user-group/dto"
	"github.com/eolinker/go-common/autowire"
	"reflect"
)

type IUserGroupModule interface {
	Get(ctx context.Context, id string) (*user_group_dto.UserGroup, error)
	List(ctx context.Context) ([]*user_group_dto.UserGroup, error)
	Create(ctx context.Context, id string, input *user_group_dto.Create) error
	Edit(ctx context.Context, id string, input *user_group_dto.Edit) error
	Delete(ctx context.Context, id string) error
	Simple(ctx context.Context) ([]*user_group_dto.Simple, error)

	AddMember(ctx context.Context, id string, member *user_group_dto.AddMember) error
	RemoveMember(ctx context.Context, id string, uid string) error
}

func init() {
	autowire.Auto[IUserGroupModule](func() reflect.Value {
		return reflect.ValueOf(new(imlUserGroupModule))
	})
}
