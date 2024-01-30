package user_group

import "time"

type UserGroup struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Creator    string    `json:"creator"`
	CreateTime time.Time `json:"createTime"`
}

type Member struct {
	UserId  string `json:"userId"`
	GroupId string `json:"groupId"`
}
