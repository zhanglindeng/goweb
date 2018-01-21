package user

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/config"
	"github.com/zhanglindeng/goweb/model"
	"github.com/zhanglindeng/goweb/model/repository"
	"github.com/zhanglindeng/goweb/validate"
	"github.com/zhanglindeng/util"
)

func Register(ctx *gin.Context) {

	ur, err := validate.UserRegisterValidate(ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	u := &model.User{
		Name:     strings.Split(ur.Email, "@")[0],
		Email:    ur.Email,
		Password: util.Md5(ur.Password), //暂时 md5
	}

	rur := repository.UserRepository{}

	if err := rur.Create(u); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		Issuer:    u.Email,
		NotBefore: time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if s, err := token.SignedString(config.AppSecret); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
	} else {
		ctx.JSON(200, gin.H{"code": 0, "token": s, "message": "ok"})
	}
}
