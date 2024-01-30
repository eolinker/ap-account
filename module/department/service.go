package department

import (
	"context"
	department_dto "gitlab.eolink.com/apinto/aoaccount/module/department/dto"
	"gitlab.eolink.com/apinto/common/autowire"
	"reflect"
)

type IDepartmentModule interface {
	CreateDepartment(ctx context.Context, department *department_dto.Create) (string, error)
	EditDepartment(ctx context.Context, id string, department *department_dto.Edit) error
	Delete(ctx context.Context, id string) error
	Tree(ctx context.Context) (*department_dto.Department, int, error)
	Simple(ctx context.Context) (*department_dto.Simple, error)
	AddMember(ctx context.Context, id string, member *department_dto.AddMember) error
	RemoveMember(ctx context.Context, id string, uid string) error
	RemoveMembers(ctx context.Context, id string, members *department_dto.RemoveMember) error
}

func init() {
	autowire.Auto[IDepartmentModule](func() reflect.Value {
		return reflect.ValueOf(new(imlDepartmentModule))
	})
}
