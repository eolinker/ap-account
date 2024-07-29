package role

import (
	"context"
	"time"

	"github.com/eolinker/go-common/auto"

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
	now := time.Now()
	r := &store.Role{
		UUID:        input.Id,
		Name:        input.Name,
		Group:       input.Group,
		Description: input.Description,
		Permit:      input.Permit,
		CreateAt:    now,
		UpdateAt:    now,
		Default:     input.Default,
	}
	return i.store.Insert(ctx, r)
}

func (i *imlRoleService) Edit(ctx context.Context, id string, input *UpdateRole) error {
	r, err := i.store.GetByUUID(ctx, id)
	if err != nil {
		return err
	}

	if input.Name != nil {
		r.Name = *input.Name
	}
	if input.Description != nil {
		r.Description = *input.Description
	}
	if input.Group != nil {
		r.Group = *input.Group
	}
	if input.Permit != nil {
		r.Permit = *input.Permit
	}
	if input.Default != nil {
		r.Default = *input.Default
	}
	r.UpdateAt = time.Now()

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

func (i *imlRoleService) OnComplete() {
	auto.RegisterService("role", i)
}

var _ IRoleMemberService = (*iMemberService)(nil)

type iMemberService struct {
	store store.IRoleMemberStore `autowired:""`
}

func (i *iMemberService) Add(ctx context.Context, input *AddMember) error {
	return i.store.Save(ctx, &store.RoleMember{
		Target: input.Target,
		Role:   input.Role,
		User:   input.User,
	})
}

func (i *iMemberService) RemoveUserRole(ctx context.Context, user string, target string) error {
	_, err := i.store.DeleteWhere(ctx, map[string]interface{}{
		"target": target,
		"user":   user,
	})
	return err
}

func (i *iMemberService) ListByTarget(ctx context.Context, target string) ([]*Member, error) {
	list, err := i.store.List(ctx, map[string]interface{}{
		"target": target,
	})
	if err != nil {
		return nil, err
	}
	return utils.SliceToSlice(list, func(e *store.RoleMember) *Member {
		return &Member{
			Role:   e.Role,
			User:   e.User,
			Target: e.Target,
		}
	}), nil
}
