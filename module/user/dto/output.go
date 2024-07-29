package user_dto

import (
	"github.com/eolinker/ap-account/service/user"
	"github.com/eolinker/go-common/auto"
)

type UserInfo struct {
	Uid        string       `json:"id"`
	Name       string       `json:"name"`
	Email      string       `json:"email"`
	Department []auto.Label `json:"department" aolabel:"department"`
	Enable     bool         `json:"enable"`
	UserRoles  []auto.Label `json:"user_roles" aolabel:"role"`
}

type UserSimple struct {
	Uid        string       `json:"id"`
	Name       string       `json:"name"`
	Email      string       `json:"email"`
	Department []auto.Label `json:"department" aolabel:"department"`
	UserRoles  []auto.Label `json:"user_roles" aolabel:"role"`
}

func CreateUserInfoFromModel(m *user.User) *UserInfo {
	dto := &UserInfo{
		Uid:        m.UID,
		Name:       m.Username,
		Email:      m.Email,
		Department: nil,
		Enable:     true,
		UserRoles:  nil,
	}
	if m.Status != 1 {
		dto.Enable = false
	}
	return dto
}
