package account

import "time"

type UserInfo struct {
	UID        string    `json:"uid"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Mobile     string    `json:"mobile"`
	NickName   string    `json:"nickname"`
	Gender     Gender    `json:"gender"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	PushToken  string    `json:"push_token"`
}

type UserItem struct {
	UID    string `json:"uid"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
