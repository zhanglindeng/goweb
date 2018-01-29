package middleware

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/cache"
	"github.com/zhanglindeng/goweb/config"
	log2 "github.com/zhanglindeng/goweb/log"
	"github.com/zhanglindeng/goweb/util"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenString := ctx.Request.Header.Get("Authorization")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid auth token")
			}
			return config.AppSecret, nil
		})

		if err != nil {
			go jwtAuthLog("parse token error, token:" + err.Error())
			ctx.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		} else {
			if !token.Valid {
				go jwtAuthLog("token invalid, token:" + tokenString)
				ctx.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			}
		}

		c := token.Claims.(jwt.MapClaims)
		email := c["iss"].(string)

		tokenMd5, err := cache.GetUserTokenHash(email)
		if err != nil {
			go jwtAuthLog("token md5 cache, email:" + email + ", error:" + err.Error())
			ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		} else {
			// tokenString eq token.Raw
			if util.Md5(token.Raw) != tokenMd5 {
				go jwtAuthLog("token md5, email:" + email)
				// 必须重新登录
				ctx.AbortWithStatusJSON(401, gin.H{"error": "token expired"})
			}
		}

		// todo add user auth log

		ctx.Next()

	}
}

func jwtAuthLog(s string) {
	log2.Info("[JwtAuth]" + s)
}
