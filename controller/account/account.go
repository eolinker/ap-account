package account

import "gitlab.eolink.com/apinto/aoaccount/service"

type AccountController interface {
	LoginOut(session string) error
}
type UserNameLoginController interface {
	LoginByUsername(username string, password string) (session string, err error)
}

type UserAdminController interface {
	AddAccount(username string, email string, password string, operator service.UserId) (service.UserId, error)
}
