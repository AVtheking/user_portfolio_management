package main

import (
	"os"

	"github.com/AVtheking/user_portfolio_management/database"
	"github.com/AVtheking/user_portfolio_management/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectDB()

	router := gin.Default()

	// Define a group for API version 1
	api := router.Group("/api/v1")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})

		routes.AuthRoutes(api)
		routes.UserRoutes(api)
		routes.PortfolioRoutes(api)
		routes.AssetRouter(api)
		routes.AnalyticsRoutes(api)
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
