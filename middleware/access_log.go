package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log2 "github.com/zhanglindeng/goweb/log"
)

func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		beginTime := time.Now()

		ctx.Next()

		log2.Info(fmt.Sprintf("IP:%s,Method:%s,Path:%s,Size:%d,ContentType:%s,StatusCode:%d,ContentLength:%d,Took:%s",
			ctx.ClientIP(),
			ctx.Request.Method,
			ctx.Request.URL.Path,
			ctx.Request.ContentLength,
			ctx.ContentType(),
			ctx.Writer.Status(),
			ctx.Writer.Size(),
			time.Since(beginTime).String(),
		))
	}
}
