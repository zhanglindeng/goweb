package data

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var keyUserToken = "user_token"

func SetUserToken(ctx *gin.Context, v interface{}) {
	ctx.Set(keyUserToken, v)
}

func GetUserToken(ctx *gin.Context) (string, error) {
	if v, exist := ctx.Get(keyUserToken); exist {
		return v.(string), nil
	}

	return "", errors.New("user token not exist")
}
