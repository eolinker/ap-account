package department

import (
	"gitlab.eolink.com/apinto/aoaccount/store"
	"gitlab.eolink.com/apinto/common/utils"
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
func ToTree(entities []*store.Department) *Node {
	root := &Node{
		Id:       "root",
		Name:     "root",
		Children: nil,
	}

	all := utils.SliceToMapO(entities, func(i *store.Department) (string, *Node) {
		n := &Node{
			Id:       i.UUID,
			Name:     i.Name,
			parent:   i.Parent,
			Children: nil,
		}

		return i.UUID, n
	})

	for _, i := range all {
		if i.parent == "" {
			continue
		}
		if p, ok := all[i.parent]; ok {
			p.Children = append(p.Children, i)
		} else {
			root.Children = append(root.Children, i)
		}
	}
	return root
}

type Node struct {
	Id       string
	Name     string
	parent   string
	Children []*Node
}

func (n *Node) Find(id string) (*Node, bool) {
	if n.Id == id {
		return n, true
	}
	for _, c := range n.Children {
		if c.Id == id {
			return c, true
		}
		if c.Children != nil {
			if c, ok := c.Find(id); ok {
				return c, true
			}
		}
	}
	return nil, false
}
func (n *Node) GetChildren() []string {
	if n.Children == nil {
		return []string{n.Id}
	}

	tmp := make([][]string, 0, len(n.Children))
	count := 0
	for _, c := range n.Children {
		cl := c.GetChildren()
		if len(cl) > 0 {
			count += len(cl)
			tmp = append(tmp, cl)
		}
	}
	ids := make([]string, 0, len(n.Children)+count+1)
	ids = append(ids, n.Id)
	ids = append(ids, utils.SliceToSlice(n.Children, func(i *Node) string {
		return i.Id
	})...)
	for _, c := range tmp {
		ids = append(ids, c...)
	}
	return ids
}

type NodeWithMembers struct {
	Id       string
	Name     string
	Children []*Node
	Members  []string // members id for this node
}
