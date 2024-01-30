package role

import (
	"gitlab.eolink.com/apinto/aoaccount/store"
	"time"
)

type Role struct {
	Id         string
	Name       string
	Creator    string
	CreateTime time.Time
}

func CreateModel(o *store.Role) *Role {
	return &Role{
		Id:         o.UUID,
		Name:       o.Name,
		Creator:    o.Creator,
		CreateTime: o.CreateTime,
	}
}
