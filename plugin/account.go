package plugin

import (
	"net/http"

	"github.com/eolinker/go-common/ignore"
	"github.com/eolinker/go-common/pm3"
)

func (p *plugin) getAccountApis() []pm3.Api {
	ignore.IgnorePath("login", http.MethodGet, "/api/v1/account/login")
	ignore.IgnorePath("login", http.MethodGet, "/api/v1/account/logout")
	ignore.IgnorePath("login", http.MethodPost, "/api/v1/account/login/username")
	return []pm3.Api{
		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/account/login/username", []string{"context", "body"}, []string{}, p.accountController.Login),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/account/logout", []string{"context"}, []string{}, p.accountController.LoginOut),

		pm3.CreateApiWidthDoc(http.MethodPut, "/api/v1/account/password/reset", []string{"context", "body"}, []string{}, p.accountController.ResetPassword),

		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/account/login", []string{"context"}, []string{"status", "channel"}, p.accountController.CheckLogin),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/account/permit/system", []string{"context"}, []string{"access"}, p.accountController.PermitSystem),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/account/profile", []string{"context"}, []string{"profile"}, p.accountController.Profile),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/account/third", []string{"context"}, []string{"drivers"}, p.accountController.ThirdDrivers),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/account/third/:driver", []string{"context", "path:driver"}, []string{"driver"}, p.accountController.ThirdDriverInfo),
		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/account/third/:driver", []string{"context", "path:driver", "body"}, []string{}, p.accountController.SaveThirdDriver),
		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/account/login/:driver", []string{"context", "path:driver", "body"}, []string{}, p.accountController.ThirdLogin),
	}
}
