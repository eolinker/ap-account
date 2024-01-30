package usage

import (
	"context"
	"sync"
)

type IUserUsageService interface {
	RemoveUser(ctx context.Context, ids ...string) error
}

var (
	lock          sync.Mutex
	usageServices []IUserUsageService
)

func RegisterUser(service IUserUsageService) {
	lock.Lock()
	defer lock.Unlock()
	usageServices = append(usageServices, service)
}

func Remove(ctx context.Context, ids ...string) error {
	lock.Lock()
	hs := make([]IUserUsageService, 0, len(usageServices))
	hs = append(hs, usageServices...)
	lock.Unlock()

	for _, h := range hs {
		err := h.RemoveUser(ctx, ids...)
		if err != nil {
			return err
		}
	}
	return nil
}
