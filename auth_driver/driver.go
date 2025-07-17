package auth_driver

import (
	"context"

	"github.com/eolinker/eosc/log"

	"github.com/eolinker/eosc"
)

var (
	defaultManager = NewManager()
)

type IDriver interface {
	Name() string
	Title() string
	ThirdLogin(ctx context.Context, args map[string]string) (string, error)
	Delete(ctx context.Context, ids ...string) error
	FilterConfig(config map[string]string)
}

type Manager struct {
	drivers eosc.Untyped[string, IDriver]
}

func (m *Manager) Register(name string, d IDriver) {
	if name == "" || d == nil {
		return
	}
	log.Info("register auth driver:", name)
	m.drivers.Set(name, d)
}

func (m *Manager) GetDriver(name string) (IDriver, bool) {
	return m.drivers.Get(name)
}

func (m *Manager) Drivers() []IDriver {
	return m.drivers.List()
}

func NewManager() *Manager {
	return &Manager{
		drivers: eosc.BuildUntyped[string, IDriver](),
	}
}

func Register(name string, d IDriver) {
	defaultManager.Register(name, d)
}
func GetDriver(name string) (IDriver, bool) {
	return defaultManager.GetDriver(name)
}

func Drivers() []IDriver {
	return defaultManager.Drivers()
}
