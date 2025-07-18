package user

import (
	"time"

	"github.com/eolinker/ap-account/service"
	"github.com/eolinker/ap-account/store"
)

type User struct {
	UID        string         `json:"uid"`
	Username   string         `json:"username"`
	Email      string         `json:"email"`
	Mobile     string         `json:"mobile"`
	Gender     service.Gender `json:"gender"`
	CreateTime time.Time      `json:"create_time"`
	UpdateTime time.Time      `json:"update_time"`
	PushToken  string         `json:"push_token"`
	From       string         `json:"from"`
	Status     int            `json:"status"`
}

func (u *User) GetLabel() string {
	return u.Username
}
func CreateModel(e *store.UserInfo) *User {
	return &User{
		UID:        e.Uid,
		Username:   e.Name,
		Email:      e.Email,
		Mobile:     e.Mobile,
		Gender:     service.Gender(e.Gender),
		CreateTime: e.CreateAt,
		UpdateTime: e.UpdateAt,
		PushToken:  e.PushToken,
		Status:     int(e.Status),
		From:       e.From,
	}
}
