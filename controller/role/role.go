package role

import (
	role_dto "github.com/eolinker/ap-account/module/role/dto"
	"github.com/eolinker/go-common/autowire"
	"github.com/gin-gonic/gin"
	"reflect"
)

type IRoleController interface {
	Get(ctx *gin.Context, id string) (*role_dto.Role, error)
	List(ctx *gin.Context) ([]*role_dto.Role, error)
	Save(ctx *gin.Context, id string, input *role_dto.Edit) error
	Create(ctx *gin.Context, id string, input *role_dto.CreateRole) error
	Delete(ctx *gin.Context, id string) error
	Simple(ctx *gin.Context, keyword string) ([]*role_dto.Simple, error)
}

func init() {
	autowire.Auto[IRoleController](func() reflect.Value {
		return reflect.ValueOf(new(imlRoleController))
	})
}
