package account

import (
	"gitlab.eolink.com/apinto/aoaccount/service/user"
	"gitlab.eolink.com/apinto/common/pm3"
)

type imlAccountModule struct {
	userService user.IUserService
	pm3.IMiddleware
}
