package user_group

import (
	"context"
	"errors"
	"gitlab.eolink.com/apinto/aoaccount/common"
	"gitlab.eolink.com/apinto/aoaccount/store"
	"gitlab.eolink.com/apinto/common/utils"
	"gorm.io/gorm"
	"time"
)

var (
	_ IUserGroupService = (*imlUserGroupService)(nil)
)

type imlUserGroupService struct {
	store store.IUserGroupStore
}

func (s *imlUserGroupService) Crete(ctx context.Context, id, name string) error {
	od, err := s.store.FirstQuery(ctx, "uuid = ? or name = ?", []interface{}{id, name}, "id")
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if od != nil {
		return errors.New("user group already exists")
	}

	return s.store.Insert(ctx, &store.UserGroup{
		Id:         0,
		UUID:       id,
		Name:       name,
		Creator:    common.UserId(ctx),
		CreateTime: time.Now(),
	})
}

func (s *imlUserGroupService) Edit(ctx context.Context, id, name string) error {
	return s.store.Transaction(ctx, func(tx context.Context) error {
		ov, err := s.store.First(tx, map[string]interface{}{"uuid": id})
		if err != nil {
			return err
		}

		du, err := s.store.FirstQuery(tx, "name <> ? and uuid = ?", []interface{}{name, id}, "id")
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {

			return err
		}
		if du != nil && du.Id != ov.Id {
			return errors.New("user group already exists")
		}

		ov.Name = name
		ov.Creator = common.UserId(ctx)
		ov.CreateTime = time.Now()

		update, err := s.store.Update(tx, ov)
		if err != nil {
			return err
		}
		if update == 0 {
			return errors.New("user group not found")
		}
		return nil

	})
}

func (s *imlUserGroupService) Delete(ctx context.Context, id string) error {
	deleteCount, err := s.store.DeleteWhere(ctx, map[string]interface{}{"uuid": id})
	if err != nil {
		return err
	}
	if deleteCount == 0 {
		return errors.New("user group not found")
	}
	return nil
}

func (s *imlUserGroupService) Get(ctx context.Context, id string) (*UserGroup, error) {
	v, e := s.store.First(ctx, map[string]interface{}{"uuid": id})
	if e != nil {
		return nil, e
	}
	return &UserGroup{
		Id:         v.UUID,
		Name:       v.Name,
		Creator:    v.Creator,
		CreateTime: v.CreateTime,
	}, nil
}

func (s *imlUserGroupService) GetList(ctx context.Context) ([]*UserGroup, error) {
	list, err := s.store.List(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return utils.SliceToSlice(list, func(e *store.UserGroup) *UserGroup {
		return &UserGroup{
			Id:         e.UUID,
			Name:       e.Name,
			Creator:    e.Creator,
			CreateTime: e.CreateTime,
		}
	}), nil
}
