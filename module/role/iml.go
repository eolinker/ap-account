package role

import (
	"context"
	"github.com/google/uuid"
	role_dto "gitlab.eolink.com/apinto/aoaccount/module/role/dto"
	"gitlab.eolink.com/apinto/aoaccount/service/role"
	"gitlab.eolink.com/apinto/common/auto"
	"gitlab.eolink.com/apinto/common/utils"
)

var (
	_ IRoleModule = (*imlRoleModule)(nil)
)

type imlRoleModule struct {
	service role.IRoleService `autowired:""`
}

func (m *imlRoleModule) Delete(ctx context.Context, id string) error {
	return m.service.Delete(ctx, id)
}

func (m *imlRoleModule) Get(ctx context.Context, id string) (*role_dto.Role, error) {
	v, err := m.service.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &role_dto.Role{
		Id:         v.Id,
		Name:       v.Name,
		Creator:    auto.UUID(v.Creator),
		CreateTime: auto.TimeLabel(v.CreateTime),
		Usage:      0,
	}, nil
}

func (m *imlRoleModule) Crete(ctx context.Context, id string, input *role_dto.CreateRole) error {
	if id == "" {
		id = uuid.NewString()
	}

	return m.service.Create(ctx, id, input.Name)
}

func (m *imlRoleModule) Edit(ctx context.Context, id string, input *role_dto.Edit) error {

	return m.service.Save(ctx, id, input.Name)
}

func (m *imlRoleModule) Simple(ctx context.Context) ([]*role_dto.Simple, error) {
	list, err := m.service.List(ctx)
	if err != nil {
		return nil, err
	}
	return utils.SliceToSlice(list, func(s *role.Role) *role_dto.Simple {
		return &role_dto.Simple{
			Id:   s.Id,
			Name: s.Name,
		}
	}), nil
}

func (m *imlRoleModule) List(ctx context.Context) ([]*role_dto.Role, error) {
	list, err := m.service.List(ctx)
	if err != nil {
		return nil, err
	}
	out := utils.SliceToSlice(list, func(s *role.Role) *role_dto.Role {
		return &role_dto.Role{
			Id:         "",
			Name:       "",
			Usage:      0,
			Creator:    auto.Label{},
			CreateTime: auto.TimeLabel{},
		}
	})
	auto.CompleteLabels(ctx, out)
	return out, nil
}
