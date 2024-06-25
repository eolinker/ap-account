package user_group_dto

import "github.com/eolinker/go-common/auto"

type UserGroup struct {
	Id         string         `json:"id"`
	Name       string         `json:"name"`
	Usage      int            `json:"usage"`
	Creator    auto.Label     `json:"creator" aolabel:"user"`
	CreateTime auto.TimeLabel `json:"create_time"`
}

type Simple struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
