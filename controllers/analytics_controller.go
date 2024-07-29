package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AVtheking/user_portfolio_management/database"
	"github.com/AVtheking/user_portfolio_management/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// GetTotalValue is a controller that returns the total value of a portfolio
func GetTotalValue(c *gin.Context) {
	portfolioId, err := strconv.Atoi(c.Param("portfolioId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid portfolio id"})
		return
	}
	portfolio := models.Portfolio{}
	result := database.DB.Model(&models.Portfolio{}).Preload("Assets").First(&portfolio, "id=?", portfolioId)
	if result.Error != nil {
		log.Println(result.Error.Error())
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": result.Error.Error()})

		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "An error occurred"})
		}
		return
	}

	totalValue := 0.0
	for _, asset := range portfolio.Assets {
		fmt.Println(asset.Value)
		totalValue += asset.Value
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": totalValue})

}

// GetAverageReturn is a controller that returns the average return of a portfolio

func GetAverageReturn(c *gin.Context) {
	portfolioID, err := strconv.Atoi(c.Param("portfolioId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid portfolio ID"})
		return
	}

	var portfolio models.Portfolio
	if err := database.DB.Model(&models.Portfolio{}).Preload("Assets").Preload("Assets.History").First(&portfolio, portfolioID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	initialTotalValue := 0.0
	finalTotalValue := 0.0

	for _, asset := range portfolio.Assets {
		if len(asset.History) > 0 {
			fmt.Println("asaset valuue", asset.History[0].Value, asset.Value)
			initialTotalValue += asset.History[0].Value
			finalTotalValue += asset.Value
		}
	}
	if initialTotalValue == 0 {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": 0.0})
		return
	}

	averageReturn := (finalTotalValue - initialTotalValue) / initialTotalValue * 100

	c.JSON(http.StatusOK, gin.H{"success": true, "data": averageReturn})
}
