package account

import (
	"context"
	"fmt"

	auth_password "github.com/eolinker/ap-account/auth_driver/auth-password"
	"github.com/eolinker/ap-account/module/account/dto"
	"github.com/eolinker/ap-account/service/user"
	"github.com/eolinker/go-common/utils"
)

var (
	_ IAccountModule = (*imlAccountModule)(nil)
)

type imlAccountModule struct {
	userService     user.IUserService          `autowired:""`
	passwordService auth_password.AuthPassword `autowired:""`
}

func (m *imlAccountModule) ResetPassword(ctx context.Context, password dto.ResetPassword) error {
	users, err := m.userService.Get(ctx, utils.UserId(ctx))
	if err != nil {
		return err
	}
	if len(users) != 1 {
		return fmt.Errorf("user not exist")
	}
	return m.passwordService.ResetPassword(ctx, users[0].Username, password.OldPassword, password.NewPassword)
}

func (m *imlAccountModule) Login(ctx context.Context, username string, password string) (string, error) {
	uid, err := m.passwordService.Login(ctx, username, password)
	if err != nil {
		// 判断是否开启访客模式，若开启，尝试访客登录
		if utils.GuestAllow() {
			return utils.GuestLogin(ctx, username, password)
		}
		return "", err
	}

	return uid, nil
}

func (m *imlAccountModule) Profile(ctx context.Context, uid string) (*dto.Profile, error) {
	// 判断是否是访客
	if utils.GuestAllow() && utils.IsGuest(ctx) {
		return &dto.Profile{
			Uid:      uid,
			Username: utils.GuestUser(),
			Email:    "",
			Phone:    "",
			Avatar:   "",
			Type:     "guest",
		}, nil
	}
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
		Type:     "user",
	}, nil
}
