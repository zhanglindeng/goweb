package route

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanglindeng/goweb/handler/api/log"
	"github.com/zhanglindeng/goweb/handler/api/menu"
	"github.com/zhanglindeng/goweb/handler/api/user"
	"github.com/zhanglindeng/goweb/middleware"
)

func api(r *gin.Engine) error {

	apiRouter := r.Group("/api")
	apiRouter.Use(middleware.JsonRequired())
	{
		userRouter := apiRouter.Group("/user")
		{
			// user register
			userRouter.POST("/register", user.Register)
			// user login
			userRouter.POST("/login", user.Login)
			// user list
			userRouter.GET("/list", user.List)
			// add user
			userRouter.POST("/add", user.Add)
		}

		logRouter := apiRouter.Group("/log")
		logRouter.Use(middleware.JwtAuth())
		{
			logRouter.GET("", log.Index)
		}

		// menu
		menuRouter := apiRouter.Group("/menu")
		{
			menuRouter.GET("", menu.Index)
			menuRouter.GET("/active", menu.Active)
			menuRouter.POST("/add", menu.Add)
			menuRouter.POST("/add/:id/submenu", menu.AddSubmenu)
			menuRouter.POST("/del/:id", menu.Del)
			menuRouter.POST("/del/:id/submenu/:sid", menu.DelSubmenu)
			menuRouter.POST("/disable/:id", menu.Disable)
			menuRouter.POST("/disable/:id/submenu/:sid", menu.DisableSubmenu)
			menuRouter.POST("/enable/:id", menu.Enable)
			menuRouter.POST("/enable/:id/submenu/:sid", menu.EnableSubmenu)
			menuRouter.POST("/update/:id", menu.Update)
			menuRouter.POST("/update/:id/submenu/:sid", menu.UpdateSubmenu)
			menuRouter.GET("/submenus/:id", menu.Submenus)
		}
	}

	return nil
}
