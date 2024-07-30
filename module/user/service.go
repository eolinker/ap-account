package user

import (
	"context"
	"reflect"

	user_dto "github.com/eolinker/ap-account/module/user/dto"
	"github.com/eolinker/go-common/autowire"
)

type IUserModule interface {
	Search(ctx context.Context, department string, keyword string) ([]*user_dto.UserInfo, error)
	Simple(ctx context.Context, keyword string) ([]*user_dto.UserSimple, error)
	AddForPassword(ctx context.Context, user *user_dto.CreateUser) (string, error)
	Disable(ctx context.Context, user *user_dto.Disable) error
	Enable(ctx context.Context, user *user_dto.Enable) error
	CountStatus(ctx context.Context, enable bool) (int, error)
	Delete(ctx context.Context, ids ...string) error
	UpdateInfo(ctx context.Context, id string, user *user_dto.EditUser) error
	UpdateUserRole(ctx context.Context, input *user_dto.UpdateUserRole) error
}

func init() {
	autowire.Auto[IUserModule](func() reflect.Value {
		return reflect.ValueOf(new(imlUserModule))
	})
}
