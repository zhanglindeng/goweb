package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/model"
	"github.com/zhanglindeng/goweb/model/repository"
)

func Create(ctx *gin.Context) {
	u := &model.User{
		Name:  "zhangsan",
		Email: "zhangsan@example.com",
	}

	ur := repository.UserRepository{}

	if err := ur.Create(u); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"User": u,
	})
}
