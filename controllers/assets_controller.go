package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AVtheking/user_portfolio_management/database"
	"github.com/AVtheking/user_portfolio_management/dto"
	"github.com/AVtheking/user_portfolio_management/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddAsset(c *gin.Context) {
	var request dto.CreateAssetRequestDto

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid data provided"})
		return
	}

	portfolioId, err := strconv.Atoi(c.Param("portfolioId"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid portfolio id"})
		return
	}
	portfolioResult := database.DB.First(&models.Portfolio{}, "id=?", portfolioId)
	if portfolioResult.Error != nil {
		log.Println(portfolioResult.Error.Error())
		if errors.Is(portfolioResult.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Portfolio not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Database error", "details": portfolioResult.Error.Error()})
		}
		return
	}

	newAsset := models.Asset{
		Name:        request.Name,
		Value:       request.Value,
		PortfolioID: uint(portfolioId),
	}

	result := database.DB.Create(&newAsset)
	if result.Error != nil {
		log.Println(result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error creating asset" + result.Error.Error()})
		return
	}

	response := dto.AssetResponseDto{
		ID:          newAsset.ID,
		Name:        newAsset.Name,
		Value:       newAsset.Value,
		PortfolioID: newAsset.PortfolioID,
		CreatedAt:   newAsset.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   newAsset.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Asset created successfully", "data": response})

}

func UpdateAsset(c *gin.Context) {
	assetId, err := strconv.Atoi(c.Param("assetId"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid asset id"})
		return
	}

	var request dto.CreateAssetRequestDto
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid data provided"})
		return
	}

	var existingAsset models.Asset
	result := database.DB.First(&existingAsset, "id = ?", assetId)
	if result.Error != nil {
		log.Println(result.Error.Error())
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Asset not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Database error", "details": result.Error.Error()})
		}
		return
	}

	assetValue := models.AssetValue{
		AssetID: existingAsset.ID,
		Value:   existingAsset.Value,
	}

	if err := database.DB.Create(&assetValue).Error; err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error updating asset" + err.Error()})
		return
	}

	existingAsset.Name = request.Name
	existingAsset.Value = request.Value
	database.DB.Save(&existingAsset)

	response := dto.AssetResponseDto{
		ID:          existingAsset.ID,
		Name:        existingAsset.Name,
		Value:       existingAsset.Value,
		PortfolioID: existingAsset.PortfolioID,
		CreatedAt:   existingAsset.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   existingAsset.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Asset updated successfully", "data": response})
}

func DeleteAsset(c *gin.Context) {
	assetId, err := strconv.Atoi(c.Param("assetId"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid asset id"})
		return
	}

	var existingAsset models.Asset
	if err := database.DB.First(&existingAsset, "id = ?", assetId).Error; err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Asset not found"})

		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Database error", "details": err.Error()})
		}
		return
	}

	result := database.DB.Delete(&existingAsset)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to delete asset", "details": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Asset deleted successfully"})
}

func GetAsset(c *gin.Context) {
	assetId, err := strconv.Atoi(c.Param("assetId"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid asset id"})
		return
	}

	var asset models.Asset
	dbError := database.DB.Model(&models.Asset{}).Preload("History").First(&asset, "id = ?", assetId).Error
	if dbError != nil {
		fmt.Println("here")
		log.Println(dbError.Error())

		if errors.Is(dbError, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Asset not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Database error", "details": dbError.Error()})
		}
		return
	}

	assetHistory := make([]dto.AssetHistory, 0)
	for _, history := range asset.History {
		assetHistory = append(assetHistory, dto.AssetHistory{

			Value: history.Value,
			Date:  history.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	response := dto.GetAssetResponseDto{
		ID:          asset.ID,
		Name:        asset.Name,
		Value:       asset.Value,
		PortfolioID: asset.PortfolioID,
		CreatedAt:   asset.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   asset.UpdatedAt.Format("2006-01-02 15:04:05"),
		History:     assetHistory,
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": response})
}

func GetAllAssets(c *gin.Context) {
	portfolioId, err := strconv.Atoi(c.Param("portfolioId"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid portfolio id"})
		return
	}

	var assets []models.Asset
	dbError := database.DB.Model(&models.Asset{}).Find(&assets, "portfolio_id=?", portfolioId).Error

	if dbError != nil {
		log.Println(dbError.Error())

		if errors.Is(dbError, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "No assets found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Database error", "details": dbError.Error()})
		}
		return
	}

	if len(assets) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "No assets found"})
		return
	}

	response := make([]dto.AssetResponseDto, 0)

	for _, asset := range assets {
		response = append(response, dto.AssetResponseDto{
			ID:          asset.ID,
			Name:        asset.Name,
			Value:       asset.Value,
			PortfolioID: asset.PortfolioID,
			CreatedAt:   asset.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   asset.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "All assets fetched successfully", "data": response})
}
