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
	Init()
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
	d, has := m.drivers.Get(name)
	if !has {
		return nil, false
	}
	d.Init()
	return d, true
}

func (m *Manager) Drivers() []IDriver {
	keys := m.drivers.Keys()
	drivers := make([]IDriver, 0, len(keys))
	for _, key := range keys {
		d, has := m.GetDriver(key)
		if !has {
			continue
		}
		drivers = append(drivers, d)
	}
	return drivers
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
