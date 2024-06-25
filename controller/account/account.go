package account

import (
	"github.com/eolinker/ap-account/middleware/login"
	"github.com/eolinker/ap-account/module/account/dto"
	"github.com/eolinker/go-common/autowire"
	"github.com/gin-gonic/gin"
	"reflect"
)

type IAccountController interface {
	login.ILoginCheck
	LoginOut(ctx *gin.Context) error
	Login(ctx *gin.Context, login *dto.Login) error

	CheckLogin(ctx *gin.Context) (string, []any, error)
	PermitSystem(ctx *gin.Context) ([]string, error)
	Profile(ctx *gin.Context) (*dto.Profile, error)
}

func init() {
	autowire.Auto[IAccountController](func() reflect.Value {
		return reflect.ValueOf(&imlAccountController{})
	})
}
