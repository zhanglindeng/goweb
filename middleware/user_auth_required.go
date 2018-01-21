package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	err2 "github.com/zhanglindeng/goweb/error"
	"github.com/zhanglindeng/goweb/middleware/data"
)

func UserAuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Authorization
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithError(http.StatusUnauthorized, err2.ErrUserUnauthorized)
		}
		log.Println("UserAuthRequired", token)

		data.SetUserToken(ctx, token)

		// TODO verify user token
		ctx.Next()
	}
}
