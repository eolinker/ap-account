package plugin

import (
	"net/http"

	"gitlab.eolink.com/apinto/common/pm3"
)

func (p *plugin) getUserGroupAPIs() []pm3.Api {
	return []pm3.Api{
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/user/groups", []string{"context"}, []string{"user_groups"}, p.userGroupController.List),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/user/group", []string{"context", "query:id"}, []string{"user_group"}, p.userGroupController.Get),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/user/group/:id", []string{"context", "path:id"}, []string{"user_group"}, p.userGroupController.Get),
		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/user/group", []string{"context", "query:id", "body"}, []string{}, p.userGroupController.Create),
		pm3.CreateApiWidthDoc(http.MethodPut, "/api/v1/user/group/:id", []string{"context", "path:id", "body"}, []string{}, p.userGroupController.Save),
		pm3.CreateApiWidthDoc(http.MethodPut, "/api/v1/user/group", []string{"context", "query:id", "body"}, []string{}, p.userGroupController.Save),
		pm3.CreateApiWidthDoc(http.MethodDelete, "/api/v1/user/group/:id", []string{"context", "path:id"}, []string{}, p.userGroupController.Delete),
		pm3.CreateApiWidthDoc(http.MethodDelete, "/api/v1/user/group", []string{"context", "query:id"}, []string{}, p.userGroupController.Delete),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/simple/user-groups", []string{"context"}, []string{"user-groups"}, p.userGroupController.Simple),

		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/user/group/member", []string{"context", "query:user_group", "body"}, []string{}, p.userGroupController.AddMember),
		pm3.CreateApiWidthDoc(http.MethodDelete, "/api/v1/user/group/member", []string{"context", "query:user_group", "query:member"}, []string{}, p.userGroupController.RemoveMember),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/user/group/members", []string{"context", "query:keyword", "query:user_group"}, []string{"members"}, p.userGroupController.Members),
	}
}
