package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/middleware/data"
	"github.com/zhanglindeng/goweb/model/repository"
)

func Info(ctx *gin.Context) {
	log.Println("user Info")

	token, err := data.GetUserToken(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Println("user token", token)

	ur := repository.UserRepository{}

	email := ctx.Query("email")

	u, err := ur.FindByEmail(email)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"User": u,
	})
}
