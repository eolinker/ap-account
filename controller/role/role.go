package role

import (
	"github.com/gin-gonic/gin"
	role_dto "gitlab.eolink.com/apinto/aoaccount/module/role/dto"
	"gitlab.eolink.com/apinto/common/autowire"
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
