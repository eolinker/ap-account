package service

import "errors"

var (
	ErrorUserNotFound = errors.New("user not found")
)

type UserId = string

type AccountService interface {
	Login(driver string, identifier string, certificate string) (UserId, error)
	AddAuth(driver string, uid string, identifier string, certificate string) error

	CheckAuth(driver string, identifier string, certificate string) (UserId, error)

	Logout(driver string, identifier string) error
	GetUserInfo(uid UserId) (UserInfo, error)
	UpdateUserInfo(uid UserId, userInfo UserInfo, operator UserId) error
	Remove(uid UserId) error
}
