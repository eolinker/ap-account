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
	UserGroups []auto.Label `json:"user_group" aolabel:"user_group"`
}

type UserSimple struct {
	Uid        string       `json:"id"`
	Name       string       `json:"name"`
	Email      string       `json:"email"`
	Department []auto.Label `json:"department" aolabel:"department"`
	UserGroups []auto.Label `json:"user_group" aolabel:"user_group"`
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
