package dto

type Login struct {
	Username string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

type ResetPassword struct {
	OldPassword string `json:"old_password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
}
