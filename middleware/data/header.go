package data

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

var keyXRequestId = "X-Request-Id"

func GetHeaderRequestId(ctx *gin.Context) string {
	requestId := ctx.Request.Header.Get(keyXRequestId)
	if len(requestId) == 0 {
		u4, err := uuid.NewV4()
		if err != nil {
			return ""
		}

		requestId = u4.String()

	}
	return requestId
}

func SetXRequestId(ctx *gin.Context, v string) {
	ctx.Set(keyXRequestId, v)
	// Set X-Request-Id header
	ctx.Writer.Header().Set(keyXRequestId, v)
}

func GetXRequestId(ctx *gin.Context) string {
	if v, exist := ctx.Get(keyXRequestId); exist {
		return v.(string)
	}

	return ""
}
