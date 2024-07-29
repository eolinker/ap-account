package plugin

import (
	"github.com/eolinker/ap-account/controller/account"
	"github.com/eolinker/ap-account/controller/department"
	"github.com/eolinker/ap-account/controller/role"
	"github.com/eolinker/ap-account/controller/users"
	"github.com/eolinker/go-common/pm3"
)

var (
	_ pm3.IPlugin           = (*plugin)(nil)
	_ pm3.IPluginMiddleware = (*plugin)(nil)
)

type plugin struct {
	apis []pm3.Api
	//userGroupController  user_group.IUserGroupController  `autowired:""`
	roleController       role.IRoleController             `autowired:""`
	departmentController department.IDepartmentController `autowired:""`
	userController       users.IUserController            `autowired:""`
	accountController    account.IAccountController       `autowired:""`
}

func (p *plugin) Middlewares() []pm3.IMiddleware {
	return []pm3.IMiddleware{
		p.accountController,
	}
}

func (p *plugin) APis() []pm3.Api {
	return p.apis
}

func (p *plugin) Name() string {
	return "users"
}
func (p *plugin) OnComplete() {
	//p.apis = append(p.apis, p.getUserGroupAPIs()...)
	p.apis = append(p.apis, p.getDepartmentApis()...)
	p.apis = append(p.apis, p.getUsersApis()...)
	p.apis = append(p.apis, p.getAccountApis()...)
	p.apis = append(p.apis, p.roleApi()...)

}
