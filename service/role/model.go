package role

import (
	"fmt"
	"time"

	"github.com/eolinker/ap-account/store"
)

type Role struct {
	Id          string
	Name        string
	Group       string
	Description string
	Permit      []string
	CreateAt    time.Time
	UpdateAt    time.Time
}

type CreateRole struct {
	Id          string
	Name        string
	Group       string
	Description string
	Permit      []string
	Default     bool
}

type UpdateRole struct {
	Name        *string
	Group       *string
	Description *string
	Permit      *[]string
	Default     *bool
}

func FromEntity(e *store.Role) *Role {
	return &Role{
		Id:          e.UUID,
		Name:        e.Name,
		Group:       e.Group,
		Description: e.Description,
		Permit:      e.Permit,
		CreateAt:    e.CreateAt,
		UpdateAt:    e.UpdateAt,
	}
}

type AddMember struct {
	Role   string
	User   string
	Target string
}

type Member struct {
	Role   string
	User   string
	Target string
}

const (
	GroupSystem = "system"
	GroupTeam   = "team"
)

func TeamTarget(id string) string {
	return fmt.Sprintf("team:%s", id)
}

func SystemTarget() string {
	return "system"
}
