package user_group

import (
	"github.com/gin-gonic/gin"
	"gitlab.eolink.com/apinto/aoaccount/module/member"
	"gitlab.eolink.com/apinto/aoaccount/module/user"
	user_group "gitlab.eolink.com/apinto/aoaccount/module/user-group"
	user_group_dto "gitlab.eolink.com/apinto/aoaccount/module/user-group/dto"
	user_dto "gitlab.eolink.com/apinto/aoaccount/module/user/dto"
)

var (
	_ IUserGroupController = (*imlUserGroupController)(nil)
)

type imlUserGroupController struct {
	module                user_group.IUserGroupModule `autowired:""`
	userModule            user.IUserModule            `autowired:""`
	userGroupMemberModule member.IMemberModule        `autowired:""`
}

func (c *imlUserGroupController) Members(ctx *gin.Context, userGroupId string) ([]*user_dto.UserInfo, error) {
	if userGroupId == "" {
		return c.userGroupMemberModule.UserGroupMember(ctx)
	} else {
		return c.userGroupMemberModule.UserGroupMember(ctx, userGroupId)
	}
}

func (c *imlUserGroupController) AddMember(ctx *gin.Context, user_group string, member *user_group_dto.AddMember) error {
	return c.module.AddMember(ctx, user_group, member)
}

func (c *imlUserGroupController) RemoveMember(ctx *gin.Context, user_group string, member string) error {
	return c.module.RemoveMember(ctx, user_group, member)
}

func (c *imlUserGroupController) Simple(ctx *gin.Context) ([]*user_group_dto.Simple, error) {
	return c.module.Simple(ctx)
}

func (c *imlUserGroupController) Get(ctx *gin.Context, id string) (*user_group_dto.UserGroup, error) {
	return c.module.Get(ctx, id)
}

func (c *imlUserGroupController) List(ctx *gin.Context) ([]*user_group_dto.UserGroup, error) {
	return c.module.List(ctx)
}

func (c *imlUserGroupController) Save(ctx *gin.Context, id string, input *user_group_dto.Edit) error {
	return c.module.Edit(ctx, id, input)
}

func (c *imlUserGroupController) Create(ctx *gin.Context, id string, input *user_group_dto.Create) error {
	return c.module.Create(ctx, id, input)
}

func (c *imlUserGroupController) Delete(ctx *gin.Context, id string) error {
	return c.module.Delete(ctx, id)
}
