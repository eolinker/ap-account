package role

import (
	"context"
	"reflect"

	"github.com/eolinker/go-common/autowire"

	"github.com/eolinker/go-common/access"

	role_dto "github.com/eolinker/ap-account/module/role/dto"
)

type IRoleModule interface {
	Add(ctx context.Context, group string, r *role_dto.CreateRole) error
	Save(ctx context.Context, group string, id string, r *role_dto.SaveRole) error
	Delete(ctx context.Context, group string, id string) error
	Get(ctx context.Context, group string, id string) (*role_dto.Role, error)
	Search(ctx context.Context, group string, keyword string) ([]*role_dto.Item, error)
	Template(ctx context.Context, group string) ([]access.Template, error)
}

func init() {
	autowire.Auto[IRoleModule](func() reflect.Value {
		return reflect.ValueOf(new(imlRoleModule))
	})
}
