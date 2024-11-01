package role

import (
	"context"
	"reflect"

	"github.com/eolinker/go-common/autowire"
)

type IRoleService interface {
	Create(ctx context.Context, input *CreateRole) error
	Edit(ctx context.Context, id string, input *UpdateRole) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*Role, error)
	SearchByGroup(ctx context.Context, keyword string, group string) ([]*Role, error)
	List(ctx context.Context, roleId ...string) ([]*Role, error)
	GetDefaultRole(ctx context.Context, group string) (*Role, error)
	GetSupperRole(ctx context.Context, group string) (*Role, error)
	ListByPermit(ctx context.Context, permit string) ([]*RoleByPermit, error)
}

type IRoleMemberService interface {
	Add(ctx context.Context, input *AddMember) error
	RemoveUserRole(ctx context.Context, target string, userId ...string) error
	List(ctx context.Context, target string, userId ...string) ([]*Member, error)
	ListByRoleIds(ctx context.Context, userId string, roleIds ...string) ([]*Member, error)
	CountByRole(ctx context.Context, target string, role string) (int64, error)
}

func init() {
	autowire.Auto[IRoleService](func() reflect.Value {
		return reflect.ValueOf(new(imlRoleService))
	})

	autowire.Auto[IRoleMemberService](func() reflect.Value {
		return reflect.ValueOf(new(iMemberService))
	})
}
