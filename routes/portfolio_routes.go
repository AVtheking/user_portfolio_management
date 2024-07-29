package routes

import (
	controller "github.com/AVtheking/user_portfolio_management/controllers"
	"github.com/AVtheking/user_portfolio_management/middlewares"
	"github.com/gin-gonic/gin"
)

// PortfolioRoutes function
func PortfolioRoutes(router *gin.RouterGroup) {
	portfolioRoutes := router.Group("/portfolio")
	portfolioRoutes.Use(middlewares.AuthMiddleWare())
	{
		portfolioRoutes.POST("/", controller.CreatePortfolio)
		portfolioRoutes.PUT("/:id", controller.UpdatePortfolio)
		portfolioRoutes.GET("/:id", controller.GetPortfolio)
		portfolioRoutes.DELETE("/:id", controller.DeletePortfolio)
	}
}
