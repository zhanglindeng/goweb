package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/middleware/data"
	"github.com/zhanglindeng/goweb/model/repository"
)

func Profile(ctx *gin.Context) {

	log.Println("user profile")

	token, err := data.GetUserToken(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Println("user token", token)

	ur := repository.UserRepository{}

	id := ctx.Query("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	u, err := ur.FindById(uint(uid))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"User": u,
	})
}
