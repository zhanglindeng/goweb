package route

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/handler"
	"github.com/zhanglindeng/goweb/helper"
	"github.com/zhanglindeng/goweb/middleware"
)

func Create() (*gin.Engine, error) {

	gin.DefaultErrorWriter = helper.SetPanicErrorWriter()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(middleware.AccessLog())

	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	corsConfig.AllowOrigins = []string{"http://localhost:4200"}
	r.Use(cors.New(corsConfig))

	r.GET("/", handler.Index)

	if err := api(r); err != nil {
		return r, err
	}

	return r, nil
}
