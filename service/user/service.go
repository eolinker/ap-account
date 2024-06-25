package user

import (
	"context"
	"github.com/eolinker/go-common/auto"
	"github.com/eolinker/go-common/autowire"
	"reflect"
)

type IUserService interface {
	Create(ctx context.Context, id string, name string, email, mobile string) (*User, error)
	SetStatus(ctx context.Context, status int, ids ...string) error
	Delete(ctx context.Context, ids ...string) error
	Search(ctx context.Context, keyword string, status int, department ...string) ([]*User, error)
	SearchUnknown(ctx context.Context, keyword string) ([]*User, error)
	Get(ctx context.Context, ids ...string) ([]*User, error)
	Update(ctx context.Context, id string, name, email, mobile *string) (*User, error)
	CountStatus(ctx context.Context, status int) (int64, error)
	auto.CompleteService
}

func init() {
	autowire.Auto[IUserService](func() reflect.Value {
		return reflect.ValueOf(new(imlUserService))
	})
}
