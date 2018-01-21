package handler

import (
	"log"
	"net/http"

	//"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/config"
	conn2 "github.com/zhanglindeng/goweb/conn"
)

func Index(ctx *gin.Context) {

	rc, err := conn2.GetRedisConn()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	//rc.Set("count", 1, persistence.DEFAULT)
	rc.Increment("count", 1)

	var count int
	rc.Get("count", &count)
	log.Println(count)

	ctx.String(http.StatusOK, config.AppName)
}
