package plugin

import (
	"github.com/eolinker/go-common/autowire"
	"github.com/eolinker/go-common/pm3"
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
