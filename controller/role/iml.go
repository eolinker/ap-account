package role

import (
	"github.com/eolinker/ap-account/module/role"
	role_dto2 "github.com/eolinker/ap-account/module/role/dto"
	"github.com/eolinker/go-common/access"
	"github.com/gin-gonic/gin"
)

var _ IRoleController = (*imlRoleController)(nil)

type imlRoleController struct {
	module role.IRoleModule `autowired:""`
}

func (i *imlRoleController) Add(ctx *gin.Context, group string, r *role_dto2.CreateRole) error {
	return i.module.Add(ctx, group, r)
}

func (i *imlRoleController) Save(ctx *gin.Context, group string, id string, r *role_dto2.SaveRole) error {
	return i.module.Save(ctx, group, id, r)
}

func (i *imlRoleController) Delete(ctx *gin.Context, group string, id string) error {
	return i.module.Delete(ctx, group, id)
}

func (i *imlRoleController) Get(ctx *gin.Context, group string, id string) (*role_dto2.Role, error) {
	return i.module.Get(ctx, group, id)
}

func (i *imlRoleController) Search(ctx *gin.Context, group string, keyword string) ([]*role_dto2.Item, error) {
	return i.module.Search(ctx, group, keyword)
}

func (i *imlRoleController) Template(ctx *gin.Context, group string) ([]access.Template, error) {
	return i.module.Template(ctx, group)
}
