package account

import "github.com/gin-gonic/gin"

type IAccountController interface {
	LoginOut(session string) error
	Login(ctx *gin.Context, username string, password string) error
}
