package department

import (
	"context"
	"gitlab.eolink.com/apinto/common/auto"
	"gitlab.eolink.com/apinto/common/autowire"
	"reflect"
)

type IDepartmentService interface {
	Create(ctx context.Context, id string, name, parent string) error
	Edit(ctx context.Context, id string, name, parent *string) error
	Get(ctx context.Context, ids ...string) ([]*Department, error)
	Tree(ctx context.Context) (*Node, error)
	Delete(ctx context.Context, id string) error
	auto.CompleteService
}

func init() {
	autowire.Auto[IDepartmentService](func() reflect.Value {
		return reflect.ValueOf(new(imlDepartmentService))
	})
}
