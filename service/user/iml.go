package user

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gitlab.eolink.com/apinto/aoaccount/service/usage"
	"gitlab.eolink.com/apinto/aoaccount/store"
	"gitlab.eolink.com/apinto/common/auto"
	"gitlab.eolink.com/apinto/common/utils"
	"gorm.io/gorm"
)

var (
	_ IUserService = (*imlUserService)(nil)
)

type imlUserService struct {
	store store.IUserInfoStore `autowired:""`
}

func (s *imlUserService) Get(ctx context.Context, ids ...string) ([]*User, error) {
	if len(ids) == 0 {
		return nil, errors.New("ids is empty")
	}
	list, err := s.store.ListQuery(ctx, "uid in (?)", []interface{}{ids}, "name asc")
	if err != nil {
		return nil, err
	}
	return utils.SliceToSlice(list, CreateModel), nil
}

func (s *imlUserService) CountStatus(ctx context.Context, status int) (int64, error) {
	return s.store.CountQuery(ctx, "status=?", status)
}

func (s *imlUserService) Create(ctx context.Context, id string, name string, email, mobile string) (*User, error) {
	if id == "" {
		id = uuid.NewString()
	}
	no := &store.UserInfo{
		Id:        0,
		Uid:       id,
		Status:    1,
		Name:      name,
		Gender:    0,
		Mobile:    mobile,
		Email:     email,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
		PushToken: email,
		IsDeleted: false,
	}
	err := s.store.Transaction(ctx, func(ctx context.Context) error {
		first, err := s.store.First(ctx, map[string]interface{}{
			"uid": id,
		})
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if first != nil {
			return errors.New("user id already exists")
		}
		return s.store.Insert(ctx, no)
	})
	if err != nil {
		return nil, err
	}
	return CreateModel(no), nil
}

func (s *imlUserService) SetStatus(ctx context.Context, status int, ids ...string) error {

	rows, err := s.store.UpdateField(ctx, "status", status, "`status` != ? and `uid` in ?", status, ids)
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("no user to update status")
	}

	return nil
}

func (s *imlUserService) Delete(ctx context.Context, ids ...string) error {
	return s.store.Transaction(ctx, func(ctx context.Context) error {
		err := s.store.SoftDeleteQuery(ctx, "uid in(?)", ids)
		if err != nil {
			return err
		}

		return usage.Remove(ctx, ids...)
	})
}

func (s *imlUserService) Search(ctx context.Context, department, keyword string) ([]*User, error) {

	where := make([]string, 0, 2)
	args := make([]interface{}, 0, 5)
	if keyword != "" {
		kv := fmt.Sprint("%", keyword, "%")
		where = append(where, "(`name` LIKE ? or `email` LIKE ? or `mobile` LIKE ? or `push_token` Like ?)")
		args = append(args, kv, kv, kv, kv)

	}
	switch strings.ToLower(department) {
	case "unknown":
		where = append(where, "not exists (select * from department_member ms where  ms.uid = user_info.uid)")
	case "disabled":
		where = append(where, "`status` != 1")
	case "":

	default:
		where = append(where, "exists (select * from department_member ms where ms.come = ? and ms.uid = user_info.uid)")
		args = append(args, department)
	}

	list, err := s.store.ListQuery(ctx, strings.Join(where, " and "), args, "`name` asc")
	if err != nil {
		return nil, err
	}
	return utils.SliceToSlice(list, CreateModel), nil
}

func (s *imlUserService) Update(ctx context.Context, id string, name, email, mobile *string) (*User, error) {
	var value *store.UserInfo
	err := s.store.Transaction(ctx, func(ctx context.Context) error {
		v, err := s.store.First(ctx, map[string]interface{}{
			"uid": id,
		})
		if err != nil {
			return err
		}
		if name != nil {
			v.Name = *name
		}
		if email != nil {
			v.Email = *email
		}
		if mobile != nil {
			v.Mobile = *email
		}
		err = s.store.Save(ctx, v)
		if err != nil {
			return err
		}
		value = v
		return nil
	})
	if err != nil {
		return nil, err
	}
	return CreateModel(value), nil
}

func (s *imlUserService) OnComplete() {
	auto.RegisterService("user", s)
}

func (s *imlUserService) GetLabels(ctx context.Context, ids ...string) map[string]string {

	if len(ids) == 0 {
		return map[string]string{}
	}
	users, err := s.store.ListQuery(ctx, "uid in (?)", []interface{}{ids}, "id asc")
	if err != nil {
		return make(map[string]string)
	}
	return utils.SliceToMapO(users, func(user *store.UserInfo) (string, string) {
		return user.Uid, user.Name
	})
}
