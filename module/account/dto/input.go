package dto

type Login struct {
	Username string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}
