package routes

import (
	controller "github.com/AVtheking/user_portfolio_management/controllers"
	"github.com/AVtheking/user_portfolio_management/middlewares"
	"github.com/gin-gonic/gin"
)

// UserRoutes function
func UserRoutes(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	userRouter.Use(middlewares.AuthMiddleWare())
	{
		userRouter.GET("/", controller.GetUser)
	}
}
