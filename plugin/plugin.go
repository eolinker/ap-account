package plugin

import "gitlab.eolink.com/apinto/common/pm3"

var (
	_ pm3.IPlugin = (*plugin)(nil)
)

type plugin struct {
	apis []pm3.Api
}

func (p *plugin) APis() []pm3.Api {
	return p.apis
}

func (p *plugin) Name() string {
	return "users"
}
func (p *plugin) OnComplete() {

}
