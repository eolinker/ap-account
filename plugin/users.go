package plugin

import (
	"gitlab.eolink.com/apinto/common/pm3"
	"net/http"
)

func (p *plugin) getUsersApis() []pm3.Api {
	return []pm3.Api{
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/user/accounts", []string{"context", "query:department", "query:keyword"}, []string{"members"}, p.userController.Search),
		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/user/account", []string{"context", "body"}, []string{"user"}, p.userController.AddForPassword),
		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/user/account/enable", []string{"context", "body"}, []string{}, p.userController.Enable),
		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/user/account/disable", []string{"context", "body"}, []string{}, p.userController.Disable),
		pm3.CreateApiWidthDoc(http.MethodDelete, "/api/v1/user/account", []string{"context", "query:id"}, []string{}, p.userController.Delete),
	}
}
