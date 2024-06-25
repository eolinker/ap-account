package plugin

import (
	"net/http"

	"github.com/eolinker/go-common/pm3"
)

func (p *plugin) getUsersApis() []pm3.Api {
	return []pm3.Api{
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/user/accounts", []string{"context", "query:department", "query:keyword"}, []string{"members"}, p.userController.Search),
		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/user/account", []string{"context", "body"}, []string{"user"}, p.userController.AddForPassword),
		pm3.CreateApiWidthDoc(http.MethodPut, "/api/v1/user/account", []string{"context", "query:id", "body"}, nil, p.userController.UpdateInfo),
		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/user/account/enable", []string{"context", "body"}, []string{}, p.userController.Enable),
		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/user/account/disable", []string{"context", "body"}, []string{}, p.userController.Disable),
		pm3.CreateApiWidthDoc(http.MethodDelete, "/api/v1/user/account", []string{"context", "query:ids"}, []string{}, p.userController.Delete),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/simple/member", []string{"context", "query:keyword"}, []string{"members"}, p.userController.Simple),
	}
}
