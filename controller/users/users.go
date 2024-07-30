package users

import (
	"reflect"

	user_dto "github.com/eolinker/ap-account/module/user/dto"
	"github.com/eolinker/go-common/autowire"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	Search(ctx *gin.Context, department string, keyword string) ([]*user_dto.UserInfo, error)
	Simple(ctx *gin.Context, keyword string) ([]*user_dto.UserSimple, error)
	AddForPassword(ctx *gin.Context, user *user_dto.CreateUser) (string, error)
	Disable(ctx *gin.Context, user *user_dto.Disable) error
	Enable(ctx *gin.Context, user *user_dto.Enable) error
	//CountStatus(ctx *gin.Context, enable bool) (int, error)
	Delete(ctx *gin.Context, id string) error
	UpdateInfo(ctx *gin.Context, id string, user *user_dto.EditUser) error
	UpdateUserRole(ctx *gin.Context, input *user_dto.UpdateUserRole) error
}

func init() {
	autowire.Auto[IUserController](func() reflect.Value {
		return reflect.ValueOf(new(imlUserController))
	})
}
