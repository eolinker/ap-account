package account

import (
	"context"
	"errors"
	"gitlab.eolink.com/apinto/aoaccount/service/usage"
	store "gitlab.eolink.com/apinto/aoaccount/store"
	"gitlab.eolink.com/apinto/common/autowire"
	"gitlab.eolink.com/apinto/common/utils"
	"gorm.io/gorm"
	"time"
)

var (
	_ IAccountService   = (*imlAccountService)(nil)
	_ autowire.Complete = (*imlAccountService)(nil)
)

type imlAccountService struct {
	store store.IUserAuthStore `autowired:""`
}

func (s *imlAccountService) Save(ctx context.Context, driver string, uid string, identifier string, certificate string) error {
	return s.store.Transaction(ctx, func(ctx context.Context) error {
		ov, err := s.store.First(ctx, map[string]interface{}{
			"driver": driver,
			"uid":    uid,
		})
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if ov == nil {
			ov = &store.UserAuth{
				Id:          0,
				Uid:         uid,
				Driver:      driver,
				Identifier:  identifier,
				Certificate: certificate,
				CreateTime:  time.Now(),
				UpdateTime:  time.Now(),
			}
			return s.store.Insert(ctx, ov)
		}
		ov.Certificate = certificate
		ov.Identifier = identifier
		updated, err := s.store.Update(ctx, ov)
		if err != nil {
			return err
		}
		if updated == 0 {
			return errors.New("update failed")
		}
		return nil
	})
}

func (s *imlAccountService) GetForUser(ctx context.Context, id string) ([]*UserAuth, error) {
	list, err := s.store.List(ctx, map[string]interface{}{
		"uid": id,
	})
	if err != nil {
		return nil, err
	}
	return utils.SliceToSlice(list, createUser), nil
}

func (s *imlAccountService) GetIdentifier(ctx context.Context, driver string, identifier string) (*UserAuth, error) {
	v, err := s.store.First(ctx, map[string]interface{}{
		"driver":     driver,
		"identifier": identifier,
	})
	if err != nil {
		return nil, err
	}
	return createUser(v), nil
}

func (s *imlAccountService) OnRemoveUsers(ctx context.Context, ids ...string) error {
	if len(ids) == 0 {
		return nil
	}
	_, err := s.store.DeleteQuery(ctx, "`uid` in (?)", ids)
	if err != nil {
		return err
	}
	return nil
}

func (s *imlAccountService) OnComplete() {
	usage.RegisterUser(s)
}
