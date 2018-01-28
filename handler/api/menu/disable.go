package menu

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/model/repository"
)

func Disable(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	mr := &repository.MenuRepository{}
	if err := mr.Disable(id); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"code": 0, "message": "ok"})
}
