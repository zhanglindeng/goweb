package route

import (
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
	r.Use(cors.Default())

	r.GET("/", handler.Index)

	//r.POST("/register", handler.Register)

	// user group
	//userRouter := r.Group("/user", middleware.UserAuthRequired())
	//{
	//	userRouter.GET("/create", user.Create)
	//	userRouter.GET("/info", user.Info)
	//	userRouter.GET("/profile", user.Profile)
	//}

	if err := api(r); err != nil {
		return r, err
	}

	return r, nil
}
