package plugin

import (
	"gitlab.eolink.com/apinto/common/pm3"
	"net/http"
)

func (p *plugin) getRoleAPIs() []pm3.Api {
	return []pm3.Api{
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/manage/roles", []string{"context"}, []string{"roles"}, p.roleController.List),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/manage/role", []string{"context", "query:id"}, []string{"role"}, p.roleController.Get),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/manage/role/:id", []string{"context", "path:id"}, []string{"role"}, p.roleController.Get),
		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/manage/role", []string{"context", "query:id", "body"}, []string{}, p.roleController.Create),
		pm3.CreateApiWidthDoc(http.MethodPut, "/api/v1/manage/role/:id", []string{"context", "path:id", "body"}, []string{}, p.roleController.Save),
		pm3.CreateApiWidthDoc(http.MethodPut, "/api/v1/manage/role", []string{"context", "query:id", "body"}, []string{}, p.roleController.Save),
		pm3.CreateApiWidthDoc(http.MethodDelete, "/api/v1/manage/role/:id", []string{"context", "path:id"}, []string{}, p.roleController.Delete),
		pm3.CreateApiWidthDoc(http.MethodDelete, "/api/v1/manage/role", []string{"context", "query:id"}, []string{}, p.roleController.Delete),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/simple/roles", []string{"context", "query:keyword"}, []string{"roles"}, p.roleController.Simple),
	}
}
