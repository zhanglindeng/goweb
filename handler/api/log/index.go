package log

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
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

	logs, count, err := alr.Find(offset, pageSize)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"code":         0,
		"message":      "ok",
		"logs":         logs,
		"count":        count,
		"current_page": page1,
		"total_page":   math.Ceil(float64(count)/float64(pageSize)),
	})
}
