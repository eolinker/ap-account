package department

import (
	"context"
	"reflect"

	department_dto "github.com/eolinker/ap-account/module/department/dto"
	"github.com/eolinker/go-common/autowire"
)

type IDepartmentModule interface {
	CreateDepartment(ctx context.Context, department *department_dto.Create) (string, error)
	EditDepartment(ctx context.Context, id string, department *department_dto.Edit) error
	Delete(ctx context.Context, id string) error
	Tree(ctx context.Context) (*department_dto.Department, error)
	Simple(ctx context.Context) (*department_dto.Simple, error)
	AddMember(ctx context.Context, member *department_dto.AddMember) error
	RemoveMember(ctx context.Context, id string, uid string) error
	RemoveMembers(ctx context.Context, id string, members *department_dto.RemoveMember) error
}

func init() {
	autowire.Auto[IDepartmentModule](func() reflect.Value {
		return reflect.ValueOf(new(imlDepartmentModule))
	})
}
