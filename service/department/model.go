package department

import (
	"gitlab.eolink.com/apinto/aoaccount/store"
	"time"
)

type Department struct {
	Id         string
	Name       string
	ParentId   string
	CreateTime time.Time
}

func fromEntity(entity *store.Department) *Department {
	return &Department{
		Id:         entity.UUID,
		Name:       entity.Name,
		ParentId:   entity.Parent,
		CreateTime: entity.CreateTime,
	}
}

type Node struct {
	Id       string
	Name     string
	Children []*Node
}
type NodeWithMembers struct {
	Id       string
	Name     string
	Children []*Node
	Members  []string // members id for this node
}
