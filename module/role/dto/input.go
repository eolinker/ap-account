package role_dto

type CreateRole struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permits     []string `json:"permits"`
}

type SaveRole struct {
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Permits     *[]string `json:"permits"`
}

type UpdateUserRole struct {
	User  string   `json:"user"`
	Roles []string `json:"roles" aocheck:"role"`
}
