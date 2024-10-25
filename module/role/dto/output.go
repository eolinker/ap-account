package role_dto

import "github.com/eolinker/go-common/auto"

type Role struct {
	ID          string         `json:"id"`
	Name        string         `json:"name" aoi18n:""`
	Group       string         `json:"group"`
	Description string         `json:"description"`
	Permits     []string       `json:"permit"`
	CreateTime  auto.TimeLabel `json:"create_time"`
	UpdateTime  auto.TimeLabel `json:"update_time"`
}

type Item struct {
	ID         string         `json:"id"`
	Name       string         `json:"name" aoi18n:""`
	Group      string         `json:"group"`
	CreateTime auto.TimeLabel `json:"create_time"`
	UpdateTime auto.TimeLabel `json:"update_time"`
	CanDelete  bool           `json:"can_delete"`
}

type SimpleItem struct {
	ID   string `json:"id"`
	Name string `json:"name" aoi18n:""`
}
