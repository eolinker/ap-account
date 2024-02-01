package session

import (
	"context"
	"fmt"
	"github.com/eolinker/eosc/log"
	"github.com/google/uuid"
	"gitlab.eolink.com/apinto/common/autowire"
	"gitlab.eolink.com/apinto/common/cache"
	"reflect"
	"time"
)

var (
	_ ISession          = (*imlSession)(nil)
	_ autowire.Complete = (*imlSession)(nil)
)

func init() {
	autowire.Auto[ISession](func() reflect.Value {
		return reflect.ValueOf(&imlSession{})
	})
}

type ISession interface {
	CreateSession(ctx context.Context, uid string) (string, error)
	Remove(ctx context.Context, session string)
	Check(ctx context.Context, sessionKey string) (Status, string)
}
type imlSession struct {
	sessionCache cache.IKVCache[SessionData, string]
}

func (s *imlSession) Remove(ctx context.Context, session string) {
	sv, err := s.sessionCache.Get(ctx, session)
	if err != nil || sv == nil {
		log.Warn("delete session error:", err)
		sv = &SessionData{
			UID:        "unknown",
			LoginTime:  0,
			Valid:      false,
			ExpireTime: time.Now().Add(ExpireTime).Unix(),
		}
	} else {
		sv.Valid = false
		sv.ExpireTime = time.Now().Add(ExpireTime).Unix()
	}
	s.sessionCache.Set(ctx, session, sv)
}

func (s *imlSession) Check(ctx context.Context, sessionKey string) (Status, string) {
	sv, err := s.sessionCache.Get(ctx, sessionKey)
	if err != nil {
		return NotLogin, ""
	}

	if sv.ExpireTime < time.Now().Unix() {
		return Expired, ""
	}

	if sv.UID == "" {
		return NotLogin, ""
	}

	sv.ExpireTime = time.Now().Add(ExpireTime).Unix()
	_ = s.sessionCache.Set(ctx, sessionKey, sv)

	return Login, sv.UID
}

func (s *imlSession) OnComplete() {
	s.sessionCache = cache.CreateKvCache[SessionData, string](ExpireTime*2, func(k string) string {
		return fmt.Sprint("session:", k)
	})
}

func (s *imlSession) CreateSession(ctx context.Context, uid string) (string, error) {
	sessionKey := uuid.NewString()

	err := s.sessionCache.Set(ctx, sessionKey, &SessionData{
		UID:        uid,
		LoginTime:  time.Now().Unix(),
		ExpireTime: time.Now().Add(ExpireTime).Unix(),
	})
	if err != nil {
		return "", err
	}

	return sessionKey, nil

}
