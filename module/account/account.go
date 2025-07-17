package account

import (
	"context"
	"reflect"

	"github.com/eolinker/ap-account/module/account/dto"
	"github.com/eolinker/go-common/autowire"
)

type IAccountModule interface {
	Login(ctx context.Context, username string, password string) (string, error)
	ResetPassword(ctx context.Context, password dto.ResetPassword) error
	Profile(ctx context.Context, uid string) (*dto.Profile, error)

	ThirdDrivers(ctx context.Context) ([]*dto.ThirdDriverItem, error)
	ThirdDriverInfo(ctx context.Context, driver string) (*dto.ThirdDriver, error)
	SaveThirdDriver(ctx context.Context, driver string, info *dto.ThirdDriver) error
	ThirdLogin(ctx context.Context, driver string, args map[string]string) (string, error)
}

func init() {
	autowire.Auto[IAccountModule](func() reflect.Value {
		return reflect.ValueOf(new(imlAccountModule))
	})
}
