package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/model"
	"github.com/zhanglindeng/goweb/model/repository"
)

func Index(ctx *gin.Context) {

	mr := repository.MenuRepository{}

	var menus []model.Menu

	if err := mr.All(&menus); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"code": 0, "message": "ok", "menus": menus})
}
