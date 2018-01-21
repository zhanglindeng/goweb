package handler

import "github.com/gin-gonic/gin"

func Register(ctx *gin.Context) {
	ctx.String(200,"user register")
}
