package common

import (
	"context"
	"github.com/gin-gonic/gin"
)

const contextUserIdKey = "user_id:context"

func UserId(ctx context.Context) string {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		return ginCtx.GetString(contextUserIdKey)
	}

	return ctx.Value(contextUserIdKey).(string)
}
func SetUserId(ctx context.Context, userId string) context.Context {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		ginCtx.Set(contextUserIdKey, userId)
		return ginCtx
	}

	return context.WithValue(ctx, contextUserIdKey, userId)
}
