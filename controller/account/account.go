package account

import (
	"github.com/gin-gonic/gin"
	"gitlab.eolink.com/apinto/aoaccount/middleware/login"
	"gitlab.eolink.com/apinto/aoaccount/module/account/dto"
	"gitlab.eolink.com/apinto/common/autowire"
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
