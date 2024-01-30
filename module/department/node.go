package department

import (
	department_dto "gitlab.eolink.com/apinto/aoaccount/module/department/dto"
	"gitlab.eolink.com/apinto/common/utils"
)

type Node struct {
	Id       string
	Name     string
	ParentId string
	Members  utils.Set[string]
	Children []*Node
}

func (n *Node) SMembers() {
	if n.Members == nil {
		n.Members = utils.NewSet[string]()
	}
	if len(n.Children) == 0 {
		return
	}

	for _, child := range n.Children {
		child.SMembers()
		n.Members.Set(child.Members.ToList()...)
	}
	return
}

func (n *Node) toDto() *department_dto.Department {

	dto := &department_dto.Department{
		Id:       n.Id,
		Name:     n.Name,
		Children: nil,
		Number:   n.Members.Size(),
	}
	dto.Children = make([]*department_dto.Department, 0, len(n.Children))
	for _, child := range n.Children {
		dto.Children = append(dto.Children, child.toDto())
	}
	return dto
}
