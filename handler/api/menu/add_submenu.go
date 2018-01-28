package menu

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/model"
	"github.com/zhanglindeng/goweb/model/repository"
	"github.com/zhanglindeng/goweb/validate"
)

func AddSubmenu(ctx *gin.Context) {

	mas, err := validate.MenuAddSubmenuValidate(ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

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

	sr := repository.SubmenuRepository{}
	s := &model.Submenu{
		Name:   mas.Name,
		Link:   mas.Link,
		Sort:   mas.Sort,
		MenuID: id,
	}

	if err := sr.Add(s); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"code": 1, "message": "ok", "submenu": s})
}
