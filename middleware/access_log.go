package middleware

import (
	"net"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log2 "github.com/zhanglindeng/goweb/log"
	"github.com/zhanglindeng/goweb/middleware/data"
	"github.com/zhanglindeng/goweb/model"
	"github.com/zhanglindeng/goweb/model/repository"
)

func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestTime := time.Now()

		requestId := data.GetHeaderRequestId(ctx)
		data.SetXRequestId(ctx, requestId)

		ctx.Next()

		durationFormat := time.Since(requestTime).String()

		// request
		requestContentLength := ctx.Request.ContentLength
		requestContentType := ctx.ContentType()
		requestMethod := ctx.Request.Method
		requestPath := ctx.Request.URL.Path
		clientIP := ctx.ClientIP()
		userAgent := ctx.Request.UserAgent()
		referer := ctx.Request.Referer()
		requestUrl := ctx.Request.URL.String()
		queryString := ctx.Request.URL.RawQuery

		// response
		responseTime := time.Now()
		remoteAddr, _, _ := net.SplitHostPort(strings.TrimSpace(ctx.Request.RemoteAddr))
		responseContentType := ctx.Writer.Header().Get("Content-Type")
		contentLength := ctx.Writer.Size()
		statusCode := ctx.Writer.Status()

		go func() {
			alr := repository.AccessLogRepository{}
			if err := alr.Create(&model.AccessLog{
				RemoteAddr:           remoteAddr,
				ClientIP:             clientIP,
				RequestTimeFloat:     requestTime.UnixNano(),
				RequestTime:          requestTime.Unix(),
				RequestID:            requestId,
				Method:               requestMethod,
				RequestContentType:   requestContentType,
				RequestContentLength: requestContentLength,
				Url:                  requestUrl,
				Path:                 requestPath,
				QueryString:          queryString,
				Referer:              referer,
				StatusCode:           statusCode,
				ContentLength:        contentLength,
				ResponseTimeFloat:    responseTime.UnixNano(),
				ResponseTime:         responseTime.Unix(),
				ResponseContentType:  responseContentType,
				Duration:             responseTime.UnixNano() - requestTime.UnixNano(),
				DurationFormat:       durationFormat,
				UserAgent:            userAgent,
			}); err != nil {
				log2.Info("[AccessLog]" + err.Error())
			}
		}()
	}
}
