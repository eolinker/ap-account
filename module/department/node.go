package department

import (
	department_dto "github.com/eolinker/ap-account/module/department/dto"
	"github.com/eolinker/ap-account/service/department"
	"github.com/eolinker/go-common/utils"
)

func toDto(node *department.Node, members map[string][]string) *department_dto.Department {

	m := utils.NewSet(members[node.Id]...)
	dto := &department_dto.Department{
		Id:       node.Id,
		Name:     node.Name,
		Children: nil,
	}

	dto.Children = make([]*department_dto.Department, 0, len(node.Children))
	for _, child := range node.Children {
		m.Set(members[child.Id]...)
		dto.Children = append(dto.Children, toDto(child, members))
	}
	dto.Number = m.Size()
	return dto
}
