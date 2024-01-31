package auth_password

import (
	"context"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"gitlab.eolink.com/apinto/aoaccount/service/account"
	"gitlab.eolink.com/apinto/common/autowire"
	"golang.org/x/crypto/pbkdf2"
	"gorm.io/gorm"
)

const (
	DriverName = `auth_password`
	saltLen    = 16
	iterations = 1000
	keyLength  = 32
)

var (
	_                     AuthPassword = (*imlAuthPassword)(nil)
	ErrorUsernameNotExist              = errors.New("user not exist")
	ErrorInvalidPassword               = errors.New("invalid password")
)

type AuthPassword interface {
	Save(ctx context.Context, id string, identifier string, certificate string) error
	Login(ctx context.Context, identifier string, certificate string) (string, error)
}

func init() {
	autowire.Auto[AuthPassword](func() reflect.Value {
		return reflect.ValueOf(new(imlAuthPassword))
	})
}

type imlAuthPassword struct {
	accountService account.IAccountService `autowired:""`
}

func (s *imlAuthPassword) Save(ctx context.Context, id string, identifier string, certificate string) error {
	secret, err := hashSecret([]byte(certificate))
	if err != nil {
		return err
	}

	err = s.accountService.Save(ctx, DriverName, id, identifier, secret)
	if err != nil {
		return err
	}
	return nil
}

func (s *imlAuthPassword) Login(ctx context.Context, identifier string, certificate string) (string, error) {

	secret, err := hashSecret([]byte(certificate))
	if err != nil {
		return "", err
	}
	auth, err := s.accountService.GetIdentifier(ctx, DriverName, identifier)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrorUsernameNotExist
		}
		return "", err
	}
	if strings.EqualFold(secret, auth.Certificate) {
		return auth.Uid, nil
	}
	return "", ErrorInvalidPassword
}
func hashSecret(secret []byte) (string, error) {
	salt, err := generateRandomSalt(saltLen)
	if err != nil {
		return "", err
	}

	// 使用 PBKDF2 密钥派生函数
	key := pbkdf2.Key(secret, salt, iterations, keyLength, sha512.New)
	return fmt.Sprintf("$pbkdf2-sha512$i=%d,l=%d$%s$%s", iterations, keyLength, base64.RawStdEncoding.EncodeToString(salt), base64.RawStdEncoding.EncodeToString(key)), nil
}

func generateRandomSalt(length int) ([]byte, error) {
	// Create a byte slice with the specified length
	salt := make([]byte, length)

	// Use crypto/rand to fill the slice with random bytes
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	// Return the salt as a hexadecimal string
	return salt, nil
}
