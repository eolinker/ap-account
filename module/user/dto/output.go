package user_dto

import (
	"gitlab.eolink.com/apinto/aoaccount/service/user"
	"gitlab.eolink.com/apinto/common/auto"
)

type UserInfo struct {
	Uid        string       `json:"id"`
	Name       string       `json:"name"`
	Email      string       `json:"email"`
	Department []auto.Label `json:"department"`
	Enable     bool         `json:"enable"`
	UserGroups []auto.Label `json:"user_group"`
}

func CreateUserInfoFromModel(m *user.User) *UserInfo {
	dto := &UserInfo{
		Uid:        m.UID,
		Name:       m.Username,
		Email:      m.Email,
		Department: nil,
		Enable:     true,
		UserGroups: nil,
	}
	if m.Status != 1 {
		dto.Enable = false
	}
	return dto
}
