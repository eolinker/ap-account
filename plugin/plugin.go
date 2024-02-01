package plugin

import (
	"gitlab.eolink.com/apinto/aoaccount/controller/account"
	"gitlab.eolink.com/apinto/aoaccount/controller/department"
	"gitlab.eolink.com/apinto/aoaccount/controller/role"
	user_group "gitlab.eolink.com/apinto/aoaccount/controller/user-group"
	"gitlab.eolink.com/apinto/aoaccount/controller/users"
	"gitlab.eolink.com/apinto/common/pm3"
)

var (
	_ pm3.IPlugin           = (*plugin)(nil)
	_ pm3.IPluginMiddleware = (*plugin)(nil)
)

type plugin struct {
	apis                 []pm3.Api
	roleController       role.IRoleController             `autowired:""`
	userGroupController  user_group.IUserGroupController  `autowired:""`
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
	p.apis = append(p.apis, p.getRoleAPIs()...)
	p.apis = append(p.apis, p.getUserGroupAPIs()...)
	p.apis = append(p.apis, p.getDepartmentApis()...)
	p.apis = append(p.apis, p.getUsersApis()...)
	p.apis = append(p.apis, p.getAccountApis()...)

}
