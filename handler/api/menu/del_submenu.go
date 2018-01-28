package menu

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/model/repository"
)

func DelSubmenu(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	mr := &repository.MenuRepository{}
	exist, err := mr.Exist(id)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	if !exist {
		ctx.JSON(200, gin.H{"code": 1, "message": "menu not found"})
		return
	}

	// submenu id
	sidStr := ctx.Param("sid")
	sid, err := strconv.Atoi(sidStr)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	sr := &repository.SubmenuRepository{}
	if err := sr.Del(sid); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"code": 0, "message": "ok"})
}
