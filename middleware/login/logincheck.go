package login

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/eolinker/ap-account/session"
	"github.com/eolinker/go-common/autowire"
	"github.com/eolinker/go-common/ignore"
	"github.com/eolinker/go-common/utils"
	"github.com/gin-gonic/gin"
)

const MiddlewareName = "login"

const i18nCookieName = "i18next"

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
	Check(method string, path string) (bool, []gin.HandlerFunc)
	//Handler(ginCtx *gin.Context)
	Name() string
	Sort() int
}

type imlLoginCheck struct {
	session.ISession `autowired:""`
}

func (m *imlLoginCheck) Sort() int {
	return -1
}

func (m *imlLoginCheck) GetSession(ginCtx *gin.Context) (string, error) {
	cookie, err := ginCtx.Cookie(session.SessionName)
	if err != nil {

		return "", err
	}
	return cookie, nil
}

func (m *imlLoginCheck) Check(method string, path string) (bool, []gin.HandlerFunc) {
	if strings.HasPrefix(path, "/api/") {
		return true, []gin.HandlerFunc{m.Handler}
	}
	return false, nil

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
		cookie, _ := ginCtx.Cookie(i18nCookieName)
		utils.SetI18n(ginCtx, cookie)
		utils.SetUserId(ginCtx, uid)
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
