package user_group

import (
	"reflect"

	"github.com/gin-gonic/gin"
	user_group_dto "gitlab.eolink.com/apinto/aoaccount/module/user-group/dto"
	user_dto "gitlab.eolink.com/apinto/aoaccount/module/user/dto"
	"gitlab.eolink.com/apinto/common/autowire"
)

type IUserGroupController interface {
	Get(ctx *gin.Context, id string) (*user_group_dto.UserGroup, error)
	List(ctx *gin.Context) ([]*user_group_dto.UserGroup, error)
	Save(ctx *gin.Context, id string, input *user_group_dto.Edit) error
	Create(ctx *gin.Context, id string, input *user_group_dto.Create) error
	Delete(ctx *gin.Context, id string) error
	Simple(ctx *gin.Context) ([]*user_group_dto.Simple, error)

	AddMember(ctx *gin.Context, user_group string, member *user_group_dto.AddMember) error
	RemoveMember(ctx *gin.Context, user_group string, member string) error
	Members(ctx *gin.Context, keyword, user_group string) ([]*user_dto.UserInfo, error)
}

func init() {
	autowire.Auto[IUserGroupController](func() reflect.Value {
		return reflect.ValueOf(new(imlUserGroupController))
	})
}
