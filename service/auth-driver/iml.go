package auth_driver

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/eolinker/go-common/utils"

	"github.com/eolinker/ap-account/store"
)

var _ IAuthService = (*imlAuthService)(nil)

type imlAuthService struct {
	store store.IAuthDriverStore `autowired:""`
}

func (i *imlAuthService) Get(ctx context.Context, id string) (*Auth, error) {
	info, err := i.store.GetByUUID(ctx, id)
	if err != nil {
		return nil, err
	}

	return FromEntity(info), nil
}

func (i *imlAuthService) List(ctx context.Context) ([]*Auth, error) {
	list, err := i.store.List(ctx, nil)
	if err != nil {
		return nil, err
	}
	return utils.SliceToSlice(list, FromEntity), nil
}

func (i *imlAuthService) ListByStatus(ctx context.Context, enable bool) ([]*Auth, error) {
	list, err := i.store.List(ctx, map[string]interface{}{"enable": enable})
	if err != nil {
		return nil, err
	}
	return utils.SliceToSlice(list, FromEntity), nil
}

func (i *imlAuthService) Save(ctx context.Context, id string, s *Save) error {
	info, err := i.store.GetByUUID(ctx, id)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		info = &store.AuthDriver{
			Uuid: id,
		}
		if s.Config != nil {
			info.Config = *s.Config
		}
		if s.Enable != nil {
			info.Enable = *s.Enable
		}
		err = i.store.Insert(ctx, info)
		if err != nil {
			return err
		}
		return nil
	}

	if s.Config != nil {
		info.Config = *s.Config
	}
	if s.Enable != nil {
		info.Enable = *s.Enable
	}
	_, err = i.store.Update(ctx, info)
	if err != nil {
		return err
	}
	return nil

}

func (i *imlAuthService) Del(ctx context.Context, id string) error {
	return i.store.DeleteUUID(ctx, id)
}
