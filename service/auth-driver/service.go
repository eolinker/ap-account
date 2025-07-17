package auth_driver

import (
	"context"
	"reflect"

	"github.com/eolinker/go-common/autowire"
)

type IAuthService interface {
	Get(ctx context.Context, id string) (*Auth, error)
	List(ctx context.Context) ([]*Auth, error)
	ListByStatus(ctx context.Context, enable bool) ([]*Auth, error)
	Save(ctx context.Context, id string, s *Save) error
	Del(ctx context.Context, id string) error
}

func init() {
	autowire.Auto[IAuthService](func() reflect.Value {
		return reflect.ValueOf(new(imlAuthService))
	})
}
