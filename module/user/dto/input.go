package user_dto

type Enable struct {
	Users []string `json:"user_ids"`
}

type Disable struct {
	Users []string `json:"user_ids"`
}

type CreateUser struct {
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Mobile      string   `json:"mobile"`
	Departments []string `json:"department_ids"`
}

type EditUser struct {
	Name *string `json:"name"`
}

type UpdateUserRole struct {
	Roles []string `json:"roles" aocheck:"role"`
	Users []string `json:"users" aocheck:"user"`
}
