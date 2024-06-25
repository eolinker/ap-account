package account

import (
	"context"
	"fmt"
	auth_password "github.com/eolinker/ap-account/auth_driver/auth-password"
	"github.com/eolinker/ap-account/module/account/dto"
	"github.com/eolinker/ap-account/service/user"
)

var (
	_ IAccountModule = (*imlAccountModule)(nil)
)

type imlAccountModule struct {
	userService     user.IUserService          `autowired:""`
	passwordService auth_password.AuthPassword `autowired:""`
}

func (m *imlAccountModule) Login(ctx context.Context, username string, password string) (string, error) {
	uid, err := m.passwordService.Login(ctx, username, password)
	if err != nil {
		return "", err
	}

	return uid, nil
}

func (m *imlAccountModule) Profile(ctx context.Context, uid string) (*dto.Profile, error) {
	users, err := m.userService.Get(ctx, uid)
	if err != nil {
		return nil, err
	}
	if len(users) != 1 {
		return nil, fmt.Errorf("user not exist")
	}
	u := users[0]
	return &dto.Profile{
		Uid:      uid,
		Username: u.Username,
		Email:    u.Email,
		Phone:    u.Mobile,
		Avatar:   "",
	}, nil
}
