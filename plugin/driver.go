package plugin

import (
	"gitlab.eolink.com/apinto/common/autowire"
	"gitlab.eolink.com/apinto/common/pm3"
)

func init() {
	pm3.Register("users", new(Driver))
}

type Driver struct {
}

func (d *Driver) Create() (pm3.IPlugin, error) {
	p := new(plugin)
	autowire.Autowired(p)
	return p, nil
}
