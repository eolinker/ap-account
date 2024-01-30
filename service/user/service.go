package user

import (
	"context"
	"gitlab.eolink.com/apinto/common/auto"
	"gitlab.eolink.com/apinto/common/autowire"
	"reflect"
)

type IUserService interface {
	Create(ctx context.Context, id string, name string, email, mobile string) (*User, error)
	SetStatus(ctx context.Context, status int, ids ...string) error
	Delete(ctx context.Context, ids ...string) error
	Search(ctx context.Context, department, keyword string) ([]*User, error)
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
