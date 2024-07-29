package role

import (
	"context"
	"time"

	"github.com/eolinker/ap-account/store"
	"github.com/eolinker/go-common/utils"
)

var _ IRoleService = (*imlRoleService)(nil)

type imlRoleService struct {
	store store.IRoleStore `autowired:""`
}

func (i *imlRoleService) List(ctx context.Context) ([]*Role, error) {
	list, err := i.store.List(ctx, nil)
	if err != nil {
		return nil, err
	}

	return utils.SliceToSlice(list, func(e *store.Role) *Role {
		return FromEntity(e)
	}), nil
}

func (i *imlRoleService) Create(ctx context.Context, input *CreateRole) error {
	return i.store.Insert(ctx, createEntityHandler(input))
}

func (i *imlRoleService) Edit(ctx context.Context, id string, input *UpdateRole) error {
	r, err := i.store.GetByUUID(ctx, id)
	if err != nil {
		return err
	}
	updateHandler(r, input)

	return i.store.Save(ctx, r)
}

func (i *imlRoleService) Delete(ctx context.Context, id string) error {
	return i.store.DeleteUUID(ctx, id)
}

func (i *imlRoleService) Get(ctx context.Context, id string) (*Role, error) {
	r, err := i.store.GetByUUID(ctx, id)
	return FromEntity(r), err
}

func (i *imlRoleService) GetDefaultRole(ctx context.Context, group string) (*Role, error) {
	r, err := i.store.First(ctx, map[string]interface{}{
		"group":   group,
		"default": true,
	})
	if err != nil {
		return nil, err
	}

	return FromEntity(r), nil
}

func (i *imlRoleService) SearchByGroup(ctx context.Context, keyword string, group string) ([]*Role, error) {
	results, err := i.store.Search(ctx, keyword, map[string]interface{}{
		"group": group,
	})
	if err != nil {
		return nil, err
	}

	return utils.SliceToSlice(results, FromEntity), nil
}

func (i *imlRoleService) GetLabels(ctx context.Context, ids ...string) map[string]string {
	if len(ids) == 0 {
		return nil
	}
	list, err := i.store.ListQuery(ctx, "`uuid` in (?)", []interface{}{ids}, "id")
	if err != nil {
		return nil
	}
	return utils.SliceToMapO(list, func(i *store.Role) (string, string) {
		return i.UUID, i.Name
	})
}

func labelHandler(e *store.Role) []string {
	return []string{e.Name, e.UUID, e.Description}
}
func uniquestHandler(i *CreateRole) []map[string]interface{} {
	return []map[string]interface{}{{"uuid": i.Id}}
}
func createEntityHandler(i *CreateRole) *store.Role {
	now := time.Now()
	return &store.Role{
		UUID:        i.Id,
		Name:        i.Name,
		Group:       i.Group,
		Description: i.Description,
		Permit:      i.Permit,
		CreateAt:    now,
		UpdateAt:    now,
		Default:     i.Default,
	}
}
func updateHandler(e *store.Role, i *UpdateRole) {
	if i.Name != nil {
		e.Name = *i.Name
	}
	if i.Description != nil {
		e.Description = *i.Description
	}
	if i.Group != nil {
		e.Group = *i.Group
	}
	if i.Permit != nil {
		e.Permit = *i.Permit
	}
	if i.Default != nil {
		e.Default = *i.Default
	}
	e.UpdateAt = time.Now()

}
