package role

import "github.com/eolinker/eosc"

type RoleByPermit struct {
	Id          string
	Name        string
	Group       string
	Description string
	Permit      string
}

type Roles eosc.Untyped[string, *RoleByPermit]

type Manager struct {
	rolesByPermit eosc.Untyped[string, Roles]
	roles         eosc.Untyped[string, *Role]
}

var (
	manager = NewManager()
)

func NewManager() *Manager {
	return &Manager{
		rolesByPermit: eosc.BuildUntyped[string, Roles](),
		roles:         eosc.BuildUntyped[string, *Role](),
	}
}

func (m *Manager) GetRolesByPermit(permit string) ([]*RoleByPermit, bool) {
	roles, has := m.rolesByPermit.Get(permit)
	if !has {
		return nil, false
	}
	list := roles.List()
	return list, len(list) > 0
}

func (m *Manager) SetRole(role *Role) {
	for _, permit := range role.Permit {
		roles, has := m.rolesByPermit.Get(permit)
		if !has {
			roles = eosc.BuildUntyped[string, *RoleByPermit]()
		}
		roles.Set(role.Id, &RoleByPermit{
			Id:          role.Id,
			Name:        role.Name,
			Group:       role.Group,
			Description: role.Description,
			Permit:      permit,
		})
		m.rolesByPermit.Set(permit, roles)
	}
	m.roles.Set(role.Id, role)
}

func (m *Manager) DeleteRole(roleId string) {
	role, has := m.roles.Get(roleId)
	if !has {
		return
	}
	for _, permit := range role.Permit {
		roles, has := m.rolesByPermit.Get(permit)
		if !has {
			continue
		}
		roles.Del(roleId)
	}
}
