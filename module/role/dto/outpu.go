package role_dto

import "github.com/eolinker/go-common/auto"

type Role struct {
	Id         string         `json:"id"`
	Name       string         `json:"name"`
	Usage      int            `json:"usage,omitempty"`
	Creator    auto.Label     `json:"creator" aolabel:"user"`
	CreateTime auto.TimeLabel `json:"create_time"`
}

type Simple struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
