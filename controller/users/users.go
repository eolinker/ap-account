package users

import (
	"github.com/gin-gonic/gin"
	user_dto "gitlab.eolink.com/apinto/aoaccount/module/user/dto"
	"gitlab.eolink.com/apinto/common/autowire"
	"reflect"
)

type IUserController interface {
	Search(ctx *gin.Context, department string, keyword string) ([]*user_dto.UserInfo, error)
	Simple(ctx *gin.Context, keyword string) ([]*user_dto.UserSimple, error)
	AddForPassword(ctx *gin.Context, user *user_dto.CreateUser) (string, error)
	Disable(ctx *gin.Context, user *user_dto.Disable) error
	Enable(ctx *gin.Context, user *user_dto.Enable) error
	//CountStatus(ctx *gin.Context, enable bool) (int, error)
	Delete(ctx *gin.Context, id string) error
}

func init() {
	autowire.Auto[IUserController](func() reflect.Value {
		return reflect.ValueOf(new(imlUserController))
	})
}
