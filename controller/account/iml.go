package account

import (
	"github.com/eolinker/ap-account/middleware/login"
	"github.com/eolinker/ap-account/module/account"
	"github.com/eolinker/ap-account/module/account/dto"
	"github.com/eolinker/ap-account/session"
	"github.com/eolinker/go-common/access"
	"github.com/eolinker/go-common/utils"
	"github.com/gin-gonic/gin"
)

var (
	_            IAccountController = (*imlAccountController)(nil)
	loginChannel                    = []any{
		dto.Channel{
			Name:   "username",
			Config: nil,
		},
	}
)

type imlAccountController struct {
	login.ILoginCheck `autowired:""`
	accountModule     account.IAccountModule `autowired:""`
	sessionService    session.ISession       `autowired:""`
}

func (c *imlAccountController) LoginOut(ctx *gin.Context) error {
	sv, err := c.GetSession(ctx)
	if err != nil {

		return nil
	}
	c.sessionService.Remove(ctx, sv)
	ctx.SetCookie(session.SessionName, "", -1, "/", "", false, false)
	return nil
}

func (c *imlAccountController) Login(ctx *gin.Context, login *dto.Login) error {
	uid, err := c.accountModule.Login(ctx, login.Username, login.Password)
	if err != nil {
		return err
	}
	newSession, err := c.sessionService.CreateSession(ctx, uid)
	if err != nil {
		return err
	}

	ctx.SetCookie(session.SessionName, newSession, int(session.ExpireTime.Seconds()), "/", "", false, false)
	return nil
}

func (c *imlAccountController) CheckLogin(ctx *gin.Context) (string, []any, error) {
	sk, err := c.GetSession(ctx)
	if err != nil {
		return "anonymous", loginChannel, nil
	}
	status, _ := c.sessionService.Check(ctx, sk)

	if status != session.Login {
		return "anonymous", loginChannel, nil
	}
	return "authorized", loginChannel, nil
}

func (c *imlAccountController) PermitSystem(ctx *gin.Context) ([]string, error) {
	uid := utils.UserId(ctx)
	if uid == "" {
		return nil, nil
	}
	al, _ := access.Get("system")

	return utils.SliceToSlice(al, func(s access.Access) string {
		return s.Name
	}), nil
}

func (c *imlAccountController) Profile(ctx *gin.Context) (*dto.Profile, error) {
	uid := utils.UserId(ctx)
	if uid == "" {
		return nil, nil
	}
	return c.accountModule.Profile(ctx, uid)
}

func (c *imlAccountController) GetSession(ctx *gin.Context) (string, error) {
	sk, err := ctx.Cookie(session.SessionName)
	if err != nil {
		return "", err
	}
	return sk, nil
}
