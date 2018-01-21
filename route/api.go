package route

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/handler/api/user"
	"github.com/zhanglindeng/goweb/middleware"
)

func api(r *gin.Engine) error {

	apiRouter := r.Group("/api", middleware.JsonRequired())
	{
		userRouter := apiRouter.Group("/user")
		{
			// user register
			userRouter.POST("/register", user.Register)
			// user login
			userRouter.POST("/login", user.Login)
		}
	}

	return nil
}
