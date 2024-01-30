package plugin

import (
	"gitlab.eolink.com/apinto/common/pm3"
	"net/http"
)

func (p *plugin) getDepartmentApis() []pm3.Api {
	return []pm3.Api{
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/user/departments", []string{"context"}, []string{"departments"}, p.departmentController.Tree),
		//pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/user/department", []string{"context", "query:id"}, []string{"department"}, p.departmentController.Detail),
		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/user/department", []string{"context", "body"}, []string{"id"}, p.departmentController.CreateDepartment),
		pm3.CreateApiWidthDoc(http.MethodPut, "/api/v1/user/department", []string{"context", "query:id", "body"}, []string{}, p.departmentController.EditDepartment),
		pm3.CreateApiWidthDoc(http.MethodDelete, "/api/v1/user/department", []string{"context", "query:id"}, []string{}, p.departmentController.Delete),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/simple/departments", []string{"context"}, []string{"department"}, p.departmentController.Simple),

		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/user/department/member", []string{"context", "query:department", "body"}, []string{}, p.departmentController.AddMember),
		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/user/department/member/remove", []string{"context", "query:department", "query:user_id"}, []string{}, p.departmentController.RemoveMember),
		pm3.CreateApiWidthDoc(http.MethodDelete, "/api/v1/user/department/members", []string{"context", "query:department", "body"}, []string{}, p.departmentController.RemoveMembers),
	}
}
