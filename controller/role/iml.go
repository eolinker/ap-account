package role

import (
	"github.com/eolinker/ap-account/module/role"
	role_dto "github.com/eolinker/ap-account/module/role/dto"
	"github.com/gin-gonic/gin"
)

var (
	_ IRoleController = (*imlRoleController)(nil)
)

type imlRoleController struct {
	module role.IRoleModule `autowired:""`
}

func (c *imlRoleController) Get(ctx *gin.Context, id string) (*role_dto.Role, error) {
	return c.module.Get(ctx, id)
}

func (c *imlRoleController) List(ctx *gin.Context) ([]*role_dto.Role, error) {
	return c.module.List(ctx)
}

func (c *imlRoleController) Save(ctx *gin.Context, id string, input *role_dto.Edit) error {
	return c.module.Edit(ctx, id, input)
}

func (c *imlRoleController) Create(ctx *gin.Context, id string, input *role_dto.CreateRole) error {
	return c.module.Crete(ctx, id, input)
}

func (c *imlRoleController) Delete(ctx *gin.Context, id string) error {
	return c.module.Delete(ctx, id)
}

func (c *imlRoleController) Simple(ctx *gin.Context, keyword string) ([]*role_dto.Simple, error) {
	return c.module.Simple(ctx, keyword)
}
