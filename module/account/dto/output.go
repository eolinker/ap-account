package dto

type Profile struct {
	Uid      string `json:"uid"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Type     string `json:"type"`
}
type Channel struct {
	Name   string      `json:"name,omitempty"`
	Config interface{} `json:"config,omitempty"`
}
