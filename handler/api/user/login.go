package user

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/config"
	"github.com/zhanglindeng/goweb/model/repository"
	"github.com/zhanglindeng/goweb/validate"
	"github.com/zhanglindeng/util"
)

func Login(ctx *gin.Context) {

	ul, err := validate.UserLoginValidate(ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	rur := &repository.UserRepository{}

	u, err := rur.FindByEmail(ul.Email)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	if u.Password != util.Md5(ul.Password) {
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
