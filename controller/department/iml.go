package department

import (
	"github.com/eolinker/ap-account/module/department"
	department_dto "github.com/eolinker/ap-account/module/department/dto"
	"github.com/eolinker/ap-account/module/user"
	"github.com/gin-gonic/gin"
)

var (
	_ IDepartmentController = (*imlDepartmentController)(nil)
)

type imlDepartmentController struct {
	module      department.IDepartmentModule `autowired:""`
	usersModule user.IUserModule             `autowired:""`
}

func (c *imlDepartmentController) Simple(ctx *gin.Context) (*department_dto.Simple, error) {
	dpsroot, err := c.module.Simple(ctx)
	if err != nil {
		return nil, err
	}
	dpsroot.Name = "所有成员"
	dpsroot.Id = ""
	return dpsroot, nil
}

func (c *imlDepartmentController) CreateDepartment(ctx *gin.Context, department *department_dto.Create) (string, error) {
	return c.module.CreateDepartment(ctx, department)
}

func (c *imlDepartmentController) EditDepartment(ctx *gin.Context, id string, department *department_dto.Edit) error {
	return c.module.EditDepartment(ctx, id, department)
}

func (c *imlDepartmentController) Delete(ctx *gin.Context, id string) error {
	return c.module.Delete(ctx, id)
}

func (c *imlDepartmentController) Tree(ctx *gin.Context) (*department_dto.Department, error) {
	dpsroot, err := c.module.Tree(ctx)
	if err != nil {
		return nil, err
	}

	users, err := c.usersModule.Search(ctx, "unknown", "")
	if err != nil {
		return nil, err
	}
	unknown := len(users)
	dpsroot.Name = "所有成员"
	dpsroot.Id = ""
	dpsroot.Number += unknown
	disableCount, err := c.usersModule.CountStatus(ctx, false)
	if err != nil {
		return nil, err
	}
	dpsroot.Children = append(dpsroot.Children, &department_dto.Department{
		Id:       "unknown",
		Name:     "未分配",
		Children: nil,
		Number:   unknown,
	}, &department_dto.Department{
		Id:       "disable",
		Name:     "禁用",
		Children: nil,
		Number:   disableCount,
	})

	return dpsroot, nil
}

func (c *imlDepartmentController) AddMember(ctx *gin.Context, member *department_dto.AddMember) error {
	return c.module.AddMember(ctx, member)
}

func (c *imlDepartmentController) RemoveMember(ctx *gin.Context, id string, uid string) error {
	return c.module.RemoveMember(ctx, id, uid)
}

func (c *imlDepartmentController) RemoveMembers(ctx *gin.Context, id string, members *department_dto.RemoveMember) error {
	return c.module.RemoveMembers(ctx, id, members)
}
