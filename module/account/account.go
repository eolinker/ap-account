package account

import (
	"context"
	"github.com/eolinker/ap-account/module/account/dto"
	"github.com/eolinker/go-common/autowire"
	"reflect"
)

type IAccountModule interface {
	Login(ctx context.Context, username string, password string) (string, error)
	ResetPassword(ctx context.Context, password dto.ResetPassword) error
	Profile(ctx context.Context, uid string) (*dto.Profile, error)
}

func init() {
	autowire.Auto[IAccountModule](func() reflect.Value {
		return reflect.ValueOf(new(imlAccountModule))
	})
}
