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

func (i *imlRoleService) GetSupperRole(ctx context.Context, group string) (*Role, error) {
	info, err := i.store.First(ctx, map[string]interface{}{
		"group":  group,
		"supper": true,
	})
	if err != nil {
		return nil, err
	}
	return FromEntity(info), err
}

func (i *imlRoleService) List(ctx context.Context, roleId ...string) ([]*Role, error) {
	w := make(map[string]interface{})
	if len(roleId) > 0 {
		w["uuid"] = roleId
	}
	list, err := i.store.List(ctx, w)
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
		Supper:      input.Supper,
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

func (i *iMemberService) CountByRole(ctx context.Context, target string, role string) (int64, error) {
	w := map[string]interface{}{
		"target": target,
		"role":   role,
	}
	return i.store.CountWhere(ctx, w)
}

func (i *iMemberService) Add(ctx context.Context, input *AddMember) error {
	return i.store.Save(ctx, &store.RoleMember{
		Target: input.Target,
		Role:   input.Role,
		User:   input.User,
	})
}

func (i *iMemberService) RemoveUserRole(ctx context.Context, target string, userId ...string) error {
	w := map[string]interface{}{
		"target": target,
	}

	if len(userId) > 0 {
		w["user"] = userId
	}

	_, err := i.store.DeleteWhere(ctx, w)
	return err
}

func (i *iMemberService) List(ctx context.Context, target string, userIds ...string) ([]*Member, error) {
	w := map[string]interface{}{
		"target": target,
	}

	if len(userIds) > 0 {
		w["user"] = userIds
	}
	list, err := i.store.List(ctx, w)
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
