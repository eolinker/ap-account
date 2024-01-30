package user_group_dto

type Create struct {
	Name string `json:"name,omitempty"`
}

type Edit struct {
	Name string `json:"name"`
}

type AddMember struct {
	Users []string `json:"ids"`
}
