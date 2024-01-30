package account

import "gitlab.eolink.com/apinto/aoaccount/store"

type UserAuth struct {
	Uid         string `json:"uid,omitempty"`
	Driver      string `json:"driver,omitempty"`
	Identifier  string `json:"identifier,omitempty"`
	Certificate string `json:"certificate,omitempty"`
}

func createUser(e *store.UserAuth) *UserAuth {
	return &UserAuth{
		Uid:         e.Uid,
		Driver:      e.Driver,
		Identifier:  e.Identifier,
		Certificate: e.Certificate,
	}
}
