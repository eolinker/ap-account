package auth_driver

import "github.com/eolinker/ap-account/store"

type Auth struct {
	Id     string `json:"id"`
	Config string `json:"config"`
	Enable bool   `json:"enable"`
}

type Save struct {
	Config *string `json:"config"`
	Enable *bool   `json:"enable"`
}

func FromEntity(i *store.AuthDriver) *Auth {
	return &Auth{
		Id:     i.Uuid,
		Config: i.Config,
		Enable: i.Enable,
	}
}
