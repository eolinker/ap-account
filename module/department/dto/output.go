package department_dto

import "gitlab.eolink.com/apinto/common/auto"

type Department struct {
	Id       string        `json:"id"`
	Name     string        `json:"name"`
	Children []*Department `json:"children"`
	Number   int           `json:"number"`
}

type Simple struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Children []*Simple `json:"children"`
	Parent   string    `json:"-"`
}
type Member struct {
	Id              string        `json:"id"`
	Name            string        `json:"name"`
	Email           string        `json:"email"`
	Departments     []*auto.Label `json:"departments" aovalue:"department"`
	DepartmentLabel string        `json:"departments_label"`
	Enable          bool          `json:"enable"`
	UserGroups      []*auto.Label `json:"user_groups" aovalue:"user_group"`
	UserGroupLabel  string        `json:"user_groups_label"`
}
