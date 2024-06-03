package department

import (
	department_dto "gitlab.eolink.com/apinto/aoaccount/module/department/dto"
	"gitlab.eolink.com/apinto/aoaccount/service/department"
	"gitlab.eolink.com/apinto/common/utils"
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
