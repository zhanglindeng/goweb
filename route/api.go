package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/handler/api/log"
	"github.com/zhanglindeng/goweb/handler/api/user"
	"github.com/zhanglindeng/goweb/middleware"
)

func api(r *gin.Engine) error {

	apiRouter := r.Group("/api")
	apiRouter.Use(middleware.JsonRequired())
	apiRouter.Use(cors.Default())
	{
		userRouter := apiRouter.Group("/user")
		{
			// user register
			userRouter.POST("/register", user.Register)
			// user login
			userRouter.POST("/login", user.Login)
		}

		logRouter := apiRouter.Group("/log")
		logRouter.Use(middleware.JwtAuth())
		{
			logRouter.GET("/", log.Index)
		}
	}

	return nil
}
