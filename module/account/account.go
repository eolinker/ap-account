package account

import (
	"context"
	"gitlab.eolink.com/apinto/aoaccount/module/account/dto"
	"gitlab.eolink.com/apinto/common/autowire"
	"reflect"
)

type IAccountModule interface {
	Login(ctx context.Context, username string, password string) (string, error)
	Profile(ctx context.Context, uid string) (*dto.Profile, error)
}

func init() {
	autowire.Auto[IAccountModule](func() reflect.Value {
		return reflect.ValueOf(new(imlAccountModule))
	})
}
