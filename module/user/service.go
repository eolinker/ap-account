package user

import (
	"context"
	user_dto "gitlab.eolink.com/apinto/aoaccount/module/user/dto"
)

type IUserModule interface {
	Search(ctx context.Context, department string, keyword string) ([]*user_dto.UserInfo, error)
	AddForPassword(ctx context.Context, user *user_dto.CreateUser) (string, error)
	Disable(ctx context.Context, user *user_dto.Disable) error
	Enable(ctx context.Context, user *user_dto.Enable) error
	CountStatus(ctx context.Context, enable bool) (int, error)
	Delete(ctx context.Context, ids ...string) error
}
