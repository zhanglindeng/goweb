package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/config"
)

func Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, config.AppName)
}
