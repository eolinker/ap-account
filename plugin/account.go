package plugin

import (
	"gitlab.eolink.com/apinto/common/ignore"
	"gitlab.eolink.com/apinto/common/pm3"
	"net/http"
)

func (p *plugin) getAccountApis() []pm3.Api {
	ignore.IgnorePath("login", http.MethodGet, "/api/v1/account/login")
	ignore.IgnorePath("login", http.MethodGet, "/api/v1/account/logout")
	ignore.IgnorePath("login", http.MethodPost, "/api/v1/account/login/username")
	return []pm3.Api{
		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/account/login/username", []string{"context", "body"}, []string{}, p.accountController.Login),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/account/logout", []string{"context"}, []string{}, p.accountController.LoginOut),

		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/account/login", []string{"context"}, []string{"status", "channel"}, p.accountController.CheckLogin),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/account/permit/system", []string{"context"}, []string{"access"}, p.accountController.PermitSystem),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/account/profile", []string{"context"}, []string{"profile"}, p.accountController.Profile),
	}
}
