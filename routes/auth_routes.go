package routes

import (
	controller "github.com/AVtheking/user_portfolio_management/controllers"
	"github.com/gin-gonic/gin"
)

// AuthRoutes function
func AuthRoutes(router *gin.RouterGroup) {

	auth := router.Group("/auth")
	{
		auth.POST("/signup", controller.SignUp)
		auth.POST("/login", controller.Login)

	}
}
