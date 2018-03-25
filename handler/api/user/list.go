package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/model"
	"github.com/zhanglindeng/goweb/model/repository"
)

func List(ctx *gin.Context) {
	rur := &repository.UserRepository{}
	var users []model.User
	if err := rur.All(&users); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"code": 0, "message": "ok", "users": users})
}
