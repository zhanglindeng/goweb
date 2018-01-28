package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/model"
	"github.com/zhanglindeng/goweb/model/repository"
	"github.com/zhanglindeng/goweb/validate"
)

func Add(ctx *gin.Context) {

	ma, err := validate.MenuAddValidate(ctx)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	mr := repository.MenuRepository{}

	m := &model.Menu{
		Name:   ma.Name,
		NameEn: ma.NameEn,
		Icon:   ma.Icon,
		Sort:   ma.Sort,
		Remark: ma.Remark,
	}

	if err := mr.Add(m); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"code": 0, "message": "ok", "menu": m})
}
