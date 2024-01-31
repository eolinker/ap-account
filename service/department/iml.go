package department

import (
	"context"
	"errors"
	store "gitlab.eolink.com/apinto/aoaccount/store"
	"gitlab.eolink.com/apinto/common/auto"
	"gitlab.eolink.com/apinto/common/utils"
	"gorm.io/gorm"
	"time"
)

var (
	_ IDepartmentService = (*imlDepartmentService)(nil)
)

type imlDepartmentService struct {
	departmentStore store.DepartmentStore       `autowired:""`
	membersStore    store.DepartmentMemberStore `autowired:""`
}

func (s *imlDepartmentService) OnComplete() {
	auto.RegisterService("department", s)
}

func (s *imlDepartmentService) GetLabels(ctx context.Context, ids ...string) map[string]string {

	departments, err := s.departmentStore.ListQuery(ctx, "uuid in(?)", []interface{}{ids}, "id")
	if err != nil {
		return map[string]string{}
	}
	return utils.SliceToMapO(departments, func(t *store.Department) (string, string) {
		return t.UUID, t.Name
	})

}

func (s *imlDepartmentService) Delete(ctx context.Context, id string) error {

	return s.departmentStore.Transaction(ctx, func(ctx context.Context) error {
		v, err := s.departmentStore.First(ctx, map[string]interface{}{
			"uuid": id,
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("department not found")
			}
			return err
		}
		if v == nil {
			return errors.New("department not found")
		}
		err = s.membersStore.Delete(ctx, id)
		if err != nil {
			return err
		}
		_, err = s.departmentStore.Delete(ctx, v.Id)
		if err != nil {
			return err
		}
		return nil
	})

}

func (s *imlDepartmentService) Get(ctx context.Context, ids ...string) ([]*Department, error) {
	if len(ids) == 0 {
		list, err := s.departmentStore.List(ctx, map[string]interface{}{})
		if err != nil {
			return nil, err
		}
		return utils.SliceToSlice(list, fromEntity), nil
	} else {
		list, err := s.departmentStore.ListQuery(ctx, "`uuid` in (?)", []interface{}{ids}, "name asc")
		if err != nil {
			return nil, err
		}

		return utils.SliceToSlice(list, fromEntity), nil

	}
}

func (s *imlDepartmentService) Create(ctx context.Context, id string, name, parent string) error {
	ev, err := s.departmentStore.First(ctx, map[string]interface{}{
		"uuid": id,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if ev != nil {
		return errors.New("department is already exists")
	}

	err = s.departmentStore.Insert(ctx, &store.Department{
		Id:         0,
		UUID:       id,
		Name:       name,
		Parent:     parent,
		CreateTime: time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *imlDepartmentService) Edit(ctx context.Context, id string, name, parent *string) error {

	ev, err := s.departmentStore.First(ctx, map[string]interface{}{
		"uuid": id,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if ev == nil {
		return errors.New("department not found")

	}
	if name != nil {
		ev.Name = *name
	}
	if parent != nil {
		if *parent != "" {
			_, err := s.departmentStore.First(ctx, map[string]interface{}{
				"uuid": *parent,
			})
			if err != nil {
				return err
			}
		}
		ev.Parent = *parent
	}
	return s.departmentStore.Save(ctx, ev)

}
