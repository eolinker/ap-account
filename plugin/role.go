package plugin

import (
	"net/http"

	"github.com/eolinker/go-common/pm3"
)

func (p *plugin) roleApi() []pm3.Api {
	return []pm3.Api{

		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/:group/role", []string{"context", "rest:group", "query:role"}, []string{"role"}, p.roleController.Get),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/:group/roles", []string{"context", "rest:group", "query:keyword"}, []string{"roles"}, p.roleController.Search),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/:group/role/template", []string{"context", "rest:group"}, []string{"permits"}, p.roleController.Template),
	}
}
