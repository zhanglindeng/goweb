package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	err2 "github.com/zhanglindeng/goweb/error"
)

func JsonRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == "POST" && ctx.Request.ContentLength > 0 {
			if ok := strings.Contains(ctx.ContentType(), "json"); !ok {
				ctx.AbortWithError(http.StatusBadRequest, err2.ErrNotJson)
			}
		}


		ctx.Next()
	}
}
