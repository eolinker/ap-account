package role

import (
	"context"
	role_dto "github.com/eolinker/ap-account/module/role/dto"
	"github.com/eolinker/go-common/autowire"
	"reflect"
)

type IRoleModule interface {
	Crete(ctx context.Context, id string, input *role_dto.CreateRole) error
	Edit(ctx context.Context, id string, input *role_dto.Edit) error
	Simple(ctx context.Context, keyword string) ([]*role_dto.Simple, error)
	List(ctx context.Context) ([]*role_dto.Role, error)
	Get(ctx context.Context, id string) (*role_dto.Role, error)
	Delete(ctx context.Context, id string) error
}

func init() {
	autowire.Auto[IRoleModule](func() reflect.Value {
		return reflect.ValueOf(new(imlRoleModule))
	})
}
