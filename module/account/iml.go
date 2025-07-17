package account

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/eolinker/ap-account/auth_driver"
	auth_driver_service "github.com/eolinker/ap-account/service/auth-driver"

	auth_password "github.com/eolinker/ap-account/auth_driver/auth-password"
	"github.com/eolinker/ap-account/module/account/dto"
	"github.com/eolinker/ap-account/service/user"
	"github.com/eolinker/go-common/utils"
)

var (
	_ IAccountModule = (*imlAccountModule)(nil)
)

type imlAccountModule struct {
	userService       user.IUserService                `autowired:""`
	authDriverService auth_driver_service.IAuthService `autowired:""`
	passwordService   auth_password.AuthPassword       `autowired:""`
}

func (m *imlAccountModule) ThirdDrivers(ctx context.Context) ([]*dto.ThirdDriverItem, error) {
	drivers := auth_driver.Drivers()
	if len(drivers) == 0 {
		return nil, nil
	}
	items := make([]*dto.ThirdDriverItem, 0, len(drivers))
	for _, d := range drivers {
		info, err := m.ThirdDriverInfo(ctx, d.Name())
		if err != nil {
			return nil, err
		}
		d.FilterConfig(info.Config)
		items = append(items, &dto.ThirdDriverItem{
			Name:   d.Title(),
			Value:  d.Name(),
			Enable: info.Enable,
			Config: info.Config,
		})
	}
	return items, nil
}

func (m *imlAccountModule) ThirdDriverInfo(ctx context.Context, driver string) (*dto.ThirdDriver, error) {
	_, ok := auth_driver.GetDriver(driver)
	if !ok {
		return nil, fmt.Errorf("driver %s not found", driver)
	}

	info, err := m.authDriverService.Get(ctx, driver)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return &dto.ThirdDriver{
			Enable: false,
			Config: nil,
		}, nil
	}
	data := make(map[string]string)
	err = json.Unmarshal([]byte(info.Config), &data)
	if err != nil {
		return nil, err
	}

	return &dto.ThirdDriver{
		Enable: info.Enable,
		Config: data,
	}, nil
}

func (m *imlAccountModule) SaveThirdDriver(ctx context.Context, driver string, info *dto.ThirdDriver) error {
	_, ok := auth_driver.GetDriver(driver)
	if !ok {
		return fmt.Errorf("driver %s not found", driver)
	}
	cfg, err := json.Marshal(info.Config)
	if err != nil {
		return fmt.Errorf("marshal config error: %w", err)
	}
	cfgStr := string(cfg)
	return m.authDriverService.Save(ctx, driver, &auth_driver_service.Save{
		Config: &cfgStr,
		Enable: &info.Enable,
	})
}

func (m *imlAccountModule) ThirdLogin(ctx context.Context, driver string, args map[string]string) (string, error) {
	d, ok := auth_driver.GetDriver(driver)
	if !ok {
		return "", fmt.Errorf("driver %s not found", driver)
	}
	info, err := m.ThirdDriverInfo(ctx, driver)
	if err != nil {
		return "", err
	}
	if !info.Enable {
		return "", fmt.Errorf("driver %s not enable", driver)
	}
	for k, v := range info.Config {
		args[k] = v
	}

	uId, err := d.ThirdLogin(ctx, args)
	if err != nil {
		return "", err
	}
	users, err := m.userService.Get(ctx, uId)
	if err != nil {
		return "", err
	}
	u := users[0]
	if u.Status != 1 {
		return "", fmt.Errorf("user %s is not active", u.Username)
	}
	return uId, nil
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
	users, err := m.userService.Get(ctx, uid)
	if err != nil {
		return "", err
	}
	u := users[0]
	if u.Status != 1 {
		return "", fmt.Errorf("user %s is not active", username)
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
	userType := "user"
	if u.From != "self-build" && u.From != "" {
		userType = "third-user"

	}
	return &dto.Profile{
		Uid:      uid,
		Username: u.Username,
		Email:    u.Email,
		Phone:    u.Mobile,
		Avatar:   "",
		Type:     userType,
	}, nil
}
