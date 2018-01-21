package user

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/model"
	"github.com/zhanglindeng/goweb/model/repository"
	"github.com/zhanglindeng/goweb/validate"
	"github.com/zhanglindeng/util"
)

func Register(ctx *gin.Context) {

	ur, err := validate.UserRegisterValidate(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	u := &model.User{
		Name:     strings.Split(ur.Email, "@")[0],
		Email:    ur.Email,
		Password: util.Md5(ur.Password), //暂时 md5
	}

	rur := repository.UserRepository{}

	if err := rur.Create(u); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
