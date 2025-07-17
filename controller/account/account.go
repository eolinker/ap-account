package account

import (
	"reflect"

	"github.com/eolinker/ap-account/middleware/login"
	"github.com/eolinker/ap-account/module/account/dto"
	"github.com/eolinker/go-common/autowire"
	"github.com/gin-gonic/gin"
)

type IAccountController interface {
	login.ILoginCheck
	LoginOut(ctx *gin.Context) error
	Login(ctx *gin.Context, login *dto.Login) error
	ResetPassword(ctx *gin.Context, input *dto.ResetPassword) error
	CheckLogin(ctx *gin.Context) (string, []any, error)
	PermitSystem(ctx *gin.Context) ([]string, error)
	Profile(ctx *gin.Context) (*dto.Profile, error)
	ThirdDrivers(ctx *gin.Context) ([]*dto.ThirdDriverItem, error)
	ThirdDriverInfo(ctx *gin.Context, driver string) (*dto.ThirdDriver, error)
	SaveThirdDriver(ctx *gin.Context, driver string, info *dto.ThirdDriver) error
	ThirdLogin(ctx *gin.Context, driver string, args *map[string]string) error
}

func init() {
	autowire.Auto[IAccountController](func() reflect.Value {
		return reflect.ValueOf(&imlAccountController{})
	})
}
