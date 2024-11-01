package plugin

import (
	"net/http"

	"github.com/eolinker/go-common/pm3"
)

func (p *plugin) roleApi() []pm3.Api {
	return []pm3.Api{

		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/:group/role", []string{"context", "rest:group", "query:role"}, []string{"role"}, p.roleController.Get, SystemSettingsRoleView),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/:group/roles", []string{"context", "rest:group", "query:keyword"}, []string{"roles"}, p.roleController.Search, SystemSettingsRoleView),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/:group/role/template", []string{"context", "rest:group"}, []string{"permits"}, p.roleController.Template),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/simple/roles", []string{"context", "query:group"}, []string{"roles"}, p.roleController.Simple),
	}
}
