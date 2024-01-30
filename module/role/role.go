package role

import (
	"context"
	role_dto "gitlab.eolink.com/apinto/aoaccount/module/role/dto"
)

type IRoleModule interface {
	Crete(ctx context.Context, id string, input *role_dto.CreateRole) error
	Edit(ctx context.Context, id string, input *role_dto.Edit) error
	Simple(ctx context.Context) ([]*role_dto.Simple, error)
	List(ctx context.Context) ([]*role_dto.Role, error)
}
