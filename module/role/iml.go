package role

import (
	"context"
	"fmt"

	"github.com/eolinker/go-common/utils"

	"github.com/eolinker/eosc/log"

	"github.com/eolinker/go-common/server"

	"github.com/eolinker/go-common/register"

	"github.com/eolinker/go-common/auto"

	"github.com/google/uuid"

	role_dto "github.com/eolinker/ap-account/module/role/dto"
	"github.com/eolinker/ap-account/service/role"
	"github.com/eolinker/go-common/access"
)

var _ IRoleModule = (*imlRoleModule)(nil)

type imlRoleModule struct {
	roleService       role.IRoleService       `autowired:""`
	roleMemberService role.IRoleMemberService `autowired:""`
}

func (i *imlRoleModule) Simple(ctx context.Context, group string) ([]*role_dto.SimpleItem, error) {
	list, err := i.roleService.SearchByGroup(ctx, "", group)
	if err != nil {
		return nil, err
	}
	return utils.SliceToSlice(list, func(item *role.Role) *role_dto.SimpleItem {
		return &role_dto.SimpleItem{
			ID:   item.Id,
			Name: item.Name,
		}
	}), nil
}

func validPermits(group string, permits []string) error {
	p, has := access.GetPermit(group)
	if !has {
		return fmt.Errorf("group %s not found", group)
	}

	for _, permit := range permits {
		err := p.Valid(permit)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *imlRoleModule) Add(ctx context.Context, group string, r *role_dto.CreateRole) error {

	err := validPermits(group, r.Permits)
	if err != nil {
		return err
	}
	if r.ID == "" {
		r.ID = uuid.New().String()
	}
	return i.roleService.Create(ctx, &role.CreateRole{
		Id:          r.ID,
		Name:        r.Name,
		Group:       group,
		Description: r.Description,
		Permit:      r.Permits,
	})
}
func (i *imlRoleModule) Save(ctx context.Context, group string, id string, r *role_dto.SaveRole) error {
	err := validPermits(group, *r.Permits)
	if err != nil {
		return err
	}

	return i.roleService.Edit(ctx, id, &role.UpdateRole{
		Name:        r.Name,
		Description: r.Description,
		Permit:      r.Permits,
	})
}
func (i *imlRoleModule) Delete(ctx context.Context, group string, id string) error {
	return i.roleService.Delete(ctx, id)
}
func (i *imlRoleModule) Get(ctx context.Context, group string, id string) (*role_dto.Role, error) {
	info, err := i.roleService.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &role_dto.Role{
		ID:          info.Id,
		Name:        info.Name,
		Description: info.Description,
		Group:       info.Group,
		Permits:     info.Permit}, nil
}
func (i *imlRoleModule) Search(ctx context.Context, group string, keyword string) ([]*role_dto.Item, error) {
	roles, err := i.roleService.SearchByGroup(ctx, keyword, group)
	if err != nil {
		return nil, err
	}
	items := make([]*role_dto.Item, 0, len(roles))
	for _, role := range roles {
		items = append(items, &role_dto.Item{
			ID:         role.Id,
			Name:       role.Name,
			Group:      role.Group,
			CanDelete:  true,
			CreateTime: auto.TimeLabel(role.CreateAt),
			UpdateTime: auto.TimeLabel(role.UpdateAt),
		})
	}
	return items, nil
}
func (i *imlRoleModule) Template(ctx context.Context, group string) ([]access.Template, error) {
	permit, has := access.GetPermit(group)
	if !has {
		return nil, fmt.Errorf("group %s not found", group)
	}

	return permit.GetTemplate(), nil
}

func (i *imlRoleModule) OnComplete() {
	register.Handle(func(v server.Server) {
		ctx := context.Background()
		roles, err := i.roleService.List(ctx)
		if err != nil {
			log.Error("init role error: ", err.Error())
			return
		}
		roleMap := utils.SliceToMapO(roles, func(r *role.Role) (string, struct{}) {
			return r.Id, struct{}{}
		})
		defaultRoles := access.Roles()
		for group, rs := range defaultRoles {

			for _, r := range rs {
				id := fmt.Sprintf("%s.%s", group, r.Name)

				if _, has := roleMap[id]; !has {
					err = i.roleService.Create(ctx, &role.CreateRole{
						Id:          id,
						Name:        r.CName,
						Description: r.CName,
						Group:       group,
						Supper:      r.Supper,
						Permit:      r.Permits,
						Default:     r.Default,
					})
					if err != nil {
						log.Error("init role error: ", err.Error())
						continue
					}
				} else {
					err = i.roleService.Edit(ctx, id, &role.UpdateRole{
						Name:        &r.CName,
						Description: &r.CName,
						Permit:      &r.Permits,
						Default:     &r.Default,
					})
					if err != nil {
						log.Error("init role error: ", err.Error())
						continue
					}
				}
			}
		}
		// 判断admin账号是否有角色
		members, err := i.roleMemberService.List(ctx, role.GroupSystem, "admin")
		if err != nil {
			log.Error("init role error: ", err.Error())
			return
		}

		if len(members) == 0 {
			r, err := i.roleService.GetSupperRole(ctx, role.GroupSystem)
			if err != nil {
				log.Error("get supper role error: ", err.Error())
				return
			}
			err = i.roleMemberService.Add(ctx, &role.AddMember{
				Role:   r.Id,
				User:   "admin",
				Target: role.GroupSystem,
			})
			if err != nil {
				log.Error("init role error: ", err.Error())
				return
			}
		}
	})
}
