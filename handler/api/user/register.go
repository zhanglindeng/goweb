package user

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/cache"
	"github.com/zhanglindeng/goweb/config"
	"github.com/zhanglindeng/goweb/model"
	"github.com/zhanglindeng/goweb/model/repository"
	"github.com/zhanglindeng/goweb/util"
	"github.com/zhanglindeng/goweb/validate"
)

func Register(ctx *gin.Context) {

	ur, err := validate.UserRegisterValidate(ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	hashedPassword, err := util.PasswordHash(ur.Password)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	u := &model.User{
		Name:     strings.Split(ur.Email, "@")[0],
		Email:    ur.Email,
		Password: hashedPassword,
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

	s, err := token.SignedString(config.AppSecret)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
	}

	// 缓存 token 的 md5
	if err := cache.SetUserTokenHash(u.Email, util.Md5(s)); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
	}

	// todo user register log

	ctx.JSON(200, gin.H{"code": 0, "token": s, "message": "ok"})
}
