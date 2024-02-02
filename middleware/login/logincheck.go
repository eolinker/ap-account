package login

import (
	"github.com/gin-gonic/gin"
	"gitlab.eolink.com/apinto/aoaccount/common"
	"gitlab.eolink.com/apinto/aoaccount/session"
	"gitlab.eolink.com/apinto/common/autowire"
	"gitlab.eolink.com/apinto/common/ignore"
	"net/http"
	"reflect"
	"strings"
)

const MiddlewareName = "login"

var (
	_ ILoginCheck = (*imlLoginCheck)(nil)

	notLogin      = []byte(`{"code":401,"msg":"not login"}`)
	sessionExpire = []byte(`{"code":401,"msg":"not login"}`)
)

func init() {
	autowire.Auto[ILoginCheck](func() reflect.Value {
		return reflect.ValueOf(&imlLoginCheck{})
	})
}

type ILoginCheck interface {
	//pm3.IMiddleware
	GetSession(ctx *gin.Context) (string, error)
	Check(method string, path string) bool
	Handler(ginCtx *gin.Context)
	Name() string
}

type imlLoginCheck struct {
	session.ISession `autowired:""`
}

func (m *imlLoginCheck) GetSession(ginCtx *gin.Context) (string, error) {
	cookie, err := ginCtx.Cookie(session.SessionName)
	if err != nil {

		return "", err
	}
	return cookie, nil
}

func (m *imlLoginCheck) Check(method string, path string) bool {
	if strings.HasPrefix(path, "/api/") {
		return true
	}
	return false

}
func NotLoginResponse(ctx *gin.Context) {

	ctx.Data(http.StatusUnauthorized, "application/json", notLogin)
	ctx.Abort()

}

func (m *imlLoginCheck) Handler(ginCtx *gin.Context) {
	notIgnore := !ignore.IsIgnorePath(MiddlewareName, ginCtx.Request.Method, ginCtx.FullPath())
	sv, err := m.GetSession(ginCtx)
	if err != nil {
		if notIgnore {
			NotLoginResponse(ginCtx)

		}
		return
	}
	status, uid := m.ISession.Check(ginCtx, sv)
	switch status {
	case session.Login:
		common.SetUserId(ginCtx, uid)
		return
	case session.Expired:
		if notIgnore {
			ginCtx.Data(http.StatusUnauthorized, "application/json", sessionExpire)
			ginCtx.Abort()
		}

		return
	case session.NotLogin:
		if notIgnore {
			NotLoginResponse(ginCtx)
		}
		return
	}

}

func (m *imlLoginCheck) Name() string {
	return MiddlewareName
}
