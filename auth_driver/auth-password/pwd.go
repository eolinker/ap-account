package auth_password

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/eolinker/ap-account/service/account"
	"github.com/eolinker/go-common/autowire"
	"golang.org/x/crypto/pbkdf2"
	"gorm.io/gorm"
)

const (
	DriverName        = `auth_password`
	defaultSaltLen    = 16
	defaultIterations = 1000
	defaultKeyLength  = 32
)

var (
	_                       AuthPassword = (*imlAuthPassword)(nil)
	ErrorUsernameNotExist                = errors.New("user not exist")
	ErrorInvalidPassword                 = errors.New("invalid password")
	ErrorInvalidOldPassword              = errors.New("invalid old password")
)

type AuthPassword interface {
	Save(ctx context.Context, id string, identifier string, certificate string) error
	Login(ctx context.Context, identifier string, certificate string) (string, error)
	Delete(ctx context.Context, ids ...string) error
	ResetPassword(ctx context.Context, identifier string, oldPassword, newPassword string) error
}

func init() {
	autowire.Auto[AuthPassword](func() reflect.Value {
		return reflect.ValueOf(new(imlAuthPassword))
	})
}

type imlAuthPassword struct {
	accountService account.IAccountService `autowired:""`
}

func (s *imlAuthPassword) Delete(ctx context.Context, ids ...string) error {
	return s.accountService.OnRemoveUsers(ctx, ids...)
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

	auth, err := s.accountService.GetIdentifier(ctx, DriverName, identifier)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrorUsernameNotExist
		}
		return "", err
	}
	if checkPasswordHash(certificate, auth.Certificate) {
		return auth.Uid, nil
	}

	return "", ErrorInvalidPassword
}

func (s *imlAuthPassword) ResetPassword(ctx context.Context, identifier string, oldPassword, newPassword string) error {
	auth, err := s.accountService.GetIdentifier(ctx, DriverName, identifier)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrorUsernameNotExist
		}
		return err
	}
	if checkPasswordHash(oldPassword, auth.Certificate) {
		secret, err := hashSecret([]byte(newPassword))
		if err != nil {
			return err
		}
		err = s.accountService.Save(ctx, DriverName, auth.Uid, auth.Identifier, secret)
		if err != nil {
			return err
		}
		return nil
	}
	return ErrorInvalidOldPassword
}
func hashSecret(secret []byte) (string, error) {
	salt, err := generateRandomSalt(defaultSaltLen)
	if err != nil {
		return "", err
	}

	// 使用 PBKDF2 密钥派生函数
	key := pbkdf2.Key(secret, salt, defaultIterations, defaultKeyLength, sha512.New)
	return fmt.Sprintf("$pbkdf2-sha512$i=%d,l=%d$%s$%s", defaultIterations, defaultKeyLength, base64.RawStdEncoding.EncodeToString(salt), base64.RawStdEncoding.EncodeToString(key)), nil
}
func checkPasswordHash(password, hash string) bool {
	iterations, keyLength, salt, key, err := readPbkdf2Hash(hash)
	if err != nil {
		return false
	}
	secret := pbkdf2.Key([]byte(password), salt, iterations, keyLength, sha512.New)
	return bytes.Equal(key, secret)
}
func readPbkdf2Hash(hash string) (int, int, []byte, []byte, error) {
	fl := strings.FieldsFunc(hash, func(r rune) bool {
		if r == '$' || r == '=' || r == ',' || r == '-' {
			return true
		}
		return false
	})
	if len(fl) != 8 || fl[0] != "pbkdf2" {
		return 0, 0, nil, nil, errors.New("invalid hash")
	}
	if fl[1] != "sha512" {
		return 0, 0, nil, nil, errors.New("not support hash")
	}

	iterations, err := strconv.Atoi(fl[3])
	if err != nil {
		return 0, 0, nil, nil, fmt.Errorf("invalid iterations: %v", err)
	}
	keyLength, err := strconv.Atoi(fl[5])
	if err != nil {
		return 0, 0, nil, nil, fmt.Errorf("invalid key length: %v", err)
	}
	salt, err := base64.RawStdEncoding.DecodeString(fl[6])
	if err != nil {
		return 0, 0, nil, nil, fmt.Errorf("invalid salt: %v", err)
	}
	key, err := base64.RawStdEncoding.DecodeString(fl[7])
	if err != nil {
		return 0, 0, nil, nil, fmt.Errorf("invalid key: %v", err)
	}
	return iterations, keyLength, salt, key, nil
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
