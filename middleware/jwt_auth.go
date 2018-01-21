package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/core/errors"
	"github.com/zhanglindeng/goweb/config"
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
			ctx.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		} else {
			if !token.Valid {
				ctx.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			}
		}

		// TODO token.Claims["iss"] user email
		// fmt.Println(token.Claims)
		ctx.Next()

	}
}
