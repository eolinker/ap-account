package users

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.eolink.com/apinto/aoaccount/module/user"
	user_dto "gitlab.eolink.com/apinto/aoaccount/module/user/dto"
)

var _ IUserController = (*imlUserController)(nil)

type imlUserController struct {
	module user.IUserModule `autowired:""`
}

func (c *imlUserController) UpdateInfo(ctx *gin.Context, id string, user *user_dto.EditUser) error {
	return c.module.UpdateInfo(ctx, id, user)
}

func (c *imlUserController) Simple(ctx *gin.Context, keyword string) ([]*user_dto.UserSimple, error) {
	return c.module.Simple(ctx, keyword)
}

func (c *imlUserController) Search(ctx *gin.Context, department string, keyword string) ([]*user_dto.UserInfo, error) {
	return c.module.Search(ctx, department, keyword)
}

func (c *imlUserController) AddForPassword(ctx *gin.Context, user *user_dto.CreateUser) (string, error) {
	return c.module.AddForPassword(ctx, user)
}

func (c *imlUserController) Disable(ctx *gin.Context, user *user_dto.Disable) error {
	return c.module.Disable(ctx, user)
}

func (c *imlUserController) Enable(ctx *gin.Context, user *user_dto.Enable) error {
	return c.module.Enable(ctx, user)
}

func (c *imlUserController) Delete(ctx *gin.Context, idStr string) error {
	if idStr == "" {
		return nil
	}
	ids := make([]string, 0)
	err := json.Unmarshal([]byte(idStr), &ids)
	if err != nil {
		return err
	}
	return c.module.Delete(ctx, ids...)
}
