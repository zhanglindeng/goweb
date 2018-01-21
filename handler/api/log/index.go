package log

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/model"
	"github.com/zhanglindeng/goweb/model/repository"
)

func Index(ctx *gin.Context) {

	alr := &repository.AccessLogRepository{}

	page := ctx.Query("page")
	page1, err := strconv.Atoi(page)
	if err != nil {
		page1 = 1
	}

	pageSize := 15
	offset := (page1 - 1) * pageSize
	var logs []model.AccessLog

	if err := alr.Find(logs, offset, pageSize); err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"code": 0, "message": "ok", "logs": logs})
}
