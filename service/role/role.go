package role

import (
	"context"
	"github.com/eolinker/go-common/autowire"
	"reflect"
)

type IRoleService interface {
	Get(ctx context.Context, id string) (*Role, error)
	Search(ctx context.Context, keyword string) ([]*Role, error)
	Save(ctx context.Context, id string, name string) error
	Create(ctx context.Context, id string, name string) error
	Delete(ctx context.Context, id string) error
}

func init() {
	autowire.Auto[IRoleService](func() reflect.Value {
		return reflect.ValueOf(new(imlRoleService))
	})
}
