package routes

import (
	controller "github.com/AVtheking/user_portfolio_management/controllers"
	"github.com/AVtheking/user_portfolio_management/middlewares"
	"github.com/gin-gonic/gin"
)

// AssetRouter function
func AssetRouter(router *gin.RouterGroup) {
	asset := router.Group("/asset")
	asset.Use(middlewares.AuthMiddleWare())
	{
		asset.POST("/:portfolioId", controller.AddAsset)
		asset.PUT("/:assetId", controller.UpdateAsset)
		asset.DELETE("/:assetId", controller.DeleteAsset)
		asset.GET("/:assetId", controller.GetAsset)
		asset.GET("/portfolio/:portfolioId", controller.GetAllAssets)

	}
}
