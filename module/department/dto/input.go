package department_dto

type Create struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentID string `json:"parent"`
}

type Edit struct {
	Id       *string `json:"id"`
	Name     *string `json:"name"`
	ParentId *string `json:"parent_id"`
}
type AddMember struct {
	UserIds       []string `json:"user_ids"`
	DepartmentIds []string `json:"department_ids"`
}

type RemoveMember struct {
	UserIds []string `json:"user_ids"`
}
