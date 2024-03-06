package department

import (
	"reflect"

	"github.com/gin-gonic/gin"
	department_dto "gitlab.eolink.com/apinto/aoaccount/module/department/dto"
	"gitlab.eolink.com/apinto/common/autowire"
)

type IDepartmentController interface {
	CreateDepartment(ctx *gin.Context, department *department_dto.Create) (string, error)
	EditDepartment(ctx *gin.Context, id string, department *department_dto.Edit) error
	Delete(ctx *gin.Context, id string) error
	Tree(ctx *gin.Context) (*department_dto.Department, error)
	AddMember(ctx *gin.Context, member *department_dto.AddMember) error
	RemoveMember(ctx *gin.Context, id string, uid string) error
	RemoveMembers(ctx *gin.Context, id string, members *department_dto.RemoveMember) error
	Simple(ctx *gin.Context) (*department_dto.Simple, error)
}

func init() {
	autowire.Auto[IDepartmentController](func() reflect.Value {
		return reflect.ValueOf(new(imlDepartmentController))
	})
}
