package role

import (
	"context"
	"errors"
	"time"

	"gitlab.eolink.com/apinto/common/auto"

	"gitlab.eolink.com/apinto/aoaccount/common"
	"gitlab.eolink.com/apinto/aoaccount/store"
	"gitlab.eolink.com/apinto/common/utils"
	"gorm.io/gorm"
)

var (
	_ IRoleService = (*imlRoleService)(nil)
)

type imlRoleService struct {
	store store.IRoleStore `autowired:""`
}

func (s *imlRoleService) OnComplete() {
	auto.RegisterService("role", s)
}

func (s *imlRoleService) Delete(ctx context.Context, id string) error {
	dcount, err := s.store.DeleteWhere(ctx, map[string]interface{}{"uuid": id})
	if err != nil {
		return err
	}
	if dcount == 0 {
		return errors.New("role not exist")
	}
	return nil
}

func (s *imlRoleService) Get(ctx context.Context, id string) (*Role, error) {
	o, err := s.store.First(ctx, map[string]interface{}{
		"uuid": id,
	})
	if err != nil {
		return nil, err
	}
	return CreateModel(o), nil
}

func (s *imlRoleService) Search(ctx context.Context, keyword string) ([]*Role, error) {

	if keyword != "" {
		list, err := s.store.ListQuery(ctx, "`name` like ?", []interface{}{"%" + keyword + "%"}, "name asc")
		if err != nil {
			return nil, err
		}
		return utils.SliceToSlice(list, CreateModel), nil
	}
	list, err := s.store.ListQuery(ctx, "", []interface{}{}, "name asc")
	if err != nil {
		return nil, err
	}
	return utils.SliceToSlice(list, CreateModel), nil
}

func (s *imlRoleService) Save(ctx context.Context, id string, name string) error {
	operator := common.UserId(ctx)

	return s.store.Transaction(ctx, func(ctx context.Context) error {
		v, err := s.store.First(ctx, map[string]interface{}{"uuid": id})
		if err != nil {
			return err
		}
		if v.Name == name {
			return nil
		}
		_, err = s.store.FirstQuery(ctx, "`uuid` = ? or `name`!=?", []interface{}{id, name}, "id")
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("role already exist")
		}
		if err != nil {
			return err
		}
		v.Name = name
		v.Creator = operator
		v.CreateTime = time.Now()

		return s.store.Save(ctx, v)

	})

}

func (s *imlRoleService) Create(ctx context.Context, id string, name string) error {
	operator := common.UserId(ctx)
	nv := &store.Role{
		Id:         0,
		UUID:       id,
		Name:       name,
		Creator:    operator,
		CreateTime: time.Now(),
	}
	return s.store.Transaction(ctx, func(ctx context.Context) error {
		ov, err := s.store.FirstQuery(ctx, "`uuid` = ? or `name`=?", []interface{}{id, name}, "id")
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if ov != nil {
			return errors.New("role already exist")
		}

		return s.store.Insert(ctx, nv)
	})

}
func (s *imlRoleService) GetLabels(ctx context.Context, ids ...string) map[string]string {

	roles, err := s.store.ListQuery(ctx, "uuid in(?)", []interface{}{ids}, "id")
	if err != nil {
		return map[string]string{}
	}
	return utils.SliceToMapO(roles, func(t *store.Role) (string, string) {
		return t.UUID, t.Name
	})

}
