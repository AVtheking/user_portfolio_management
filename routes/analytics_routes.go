package routes

import (
	controller "github.com/AVtheking/user_portfolio_management/controllers"
	"github.com/gin-gonic/gin"
)

// AnalyticsRoutes function
func AnalyticsRoutes(router *gin.RouterGroup) {
	analytics := router.Group("/analytics")
	{
		analytics.GET("/totalValue/:portfolioId", controller.GetTotalValue)
		analytics.GET("/averageReturn/:portfolioId", controller.GetAverageReturn)
	}
}
