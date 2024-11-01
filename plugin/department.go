package plugin

import (
	"net/http"

	"github.com/eolinker/go-common/pm3"
)

func (p *plugin) getDepartmentApis() []pm3.Api {
	return []pm3.Api{
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/user/departments", []string{"context"}, []string{"departments"}, p.departmentController.Tree, SystemSettingsAccountView),
		//pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/user/department", []string{"context", "query:id"}, []string{"department"}, p.departmentController.Detail),
		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/user/department", []string{"context", "body"}, []string{"id"}, p.departmentController.CreateDepartment, SystemSettingsAccountManager),
		pm3.CreateApiWidthDoc(http.MethodPut, "/api/v1/user/department", []string{"context", "query:id", "body"}, []string{}, p.departmentController.EditDepartment, SystemSettingsAccountManager),
		pm3.CreateApiWidthDoc(http.MethodDelete, "/api/v1/user/department", []string{"context", "query:id"}, []string{}, p.departmentController.Delete, SystemSettingsAccountManager),
		pm3.CreateApiWidthDoc(http.MethodGet, "/api/v1/simple/departments", []string{"context"}, []string{"department"}, p.departmentController.Simple),

		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/user/department/member", []string{"context", "body"}, []string{}, p.departmentController.AddMember, SystemSettingsAccountManager),
		pm3.CreateApiWidthDoc(http.MethodPost, "/api/v1/user/department/member/remove", []string{"context", "query:department", "body"}, []string{}, p.departmentController.RemoveMembers, SystemSettingsAccountManager),
		pm3.CreateApiWidthDoc(http.MethodDelete, "/api/v1/user/department/members", []string{"context", "query:department", "query:user_id"}, []string{}, p.departmentController.RemoveMember, SystemSettingsAccountManager),
	}
}
