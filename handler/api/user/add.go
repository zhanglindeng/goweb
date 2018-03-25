package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/model"
	"github.com/zhanglindeng/goweb/model/repository"
	"github.com/zhanglindeng/goweb/util"
	"github.com/zhanglindeng/goweb/validate"
)

func Add(ctx *gin.Context) {
	data, err := validate.UserAddValidate(ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	hashedPassword, err := util.PasswordHash(data.Password)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})

	}

	u := &model.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: hashedPassword,
	}

	rur := repository.UserRepository{}

	if err := rur.Create(u); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	// hide password
	u.Password = ""

	ctx.JSON(200, gin.H{"code": 0, "user": u, "message": "ok"})
}
