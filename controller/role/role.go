package role

import (
	"reflect"

	role_dto2 "github.com/eolinker/ap-account/module/role/dto"

	"github.com/eolinker/go-common/access"
	"github.com/eolinker/go-common/autowire"
	"github.com/gin-gonic/gin"
)

type IRoleController interface {
	Add(ctx *gin.Context, group string, r *role_dto2.CreateRole) error
	Save(ctx *gin.Context, group string, id string, r *role_dto2.SaveRole) error
	Delete(ctx *gin.Context, group string, id string) error
	Get(ctx *gin.Context, group string, id string) (*role_dto2.Role, error)
	Search(ctx *gin.Context, group string, keyword string) ([]*role_dto2.Item, error)
	Template(ctx *gin.Context, group string) ([]access.Template, error)
	Simple(ctx *gin.Context, group string) ([]*role_dto2.SimpleItem, error)
}

func init() {
	autowire.Auto[IRoleController](func() reflect.Value {
		return reflect.ValueOf(new(imlRoleController))
	})
}
