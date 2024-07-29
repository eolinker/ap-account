package role

import (
	"context"
	"reflect"

	"github.com/eolinker/go-common/autowire"
)

type IRoleService interface {
	Create(ctx context.Context, input *CreateRole) error
	Edit(ctx context.Context, id string, input *UpdateRole) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*Role, error)
	SearchByGroup(ctx context.Context, keyword string, group string) ([]*Role, error)
	List(ctx context.Context) ([]*Role, error)
	GetDefaultRole(ctx context.Context, group string) (*Role, error)
}

func init() {
	autowire.Auto[IRoleService](func() reflect.Value {
		return reflect.ValueOf(new(imlRoleService))
	})
}
