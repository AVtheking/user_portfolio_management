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

func CreatePortfolio(c *gin.Context) {
	var portfolio dto.CreatePortfolioRequestDto

	if err := c.ShouldBindBodyWithJSON(&portfolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid data provided" + err.Error()})
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User not authenticated"})
		return
	}

	newPortfolio := models.Portfolio{
		UserID: userId.(uint),
		Name:   portfolio.Name,
	}

	createResult := database.DB.Create(&newPortfolio)
	if createResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": createResult.Error.Error()})
		return
	}

	createPortfolioResponse := dto.PortfolioResponseDto{
		ID:        newPortfolio.ID,
		UserID:    newPortfolio.UserID,
		Name:      newPortfolio.Name,
		CreatedAt: newPortfolio.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: newPortfolio.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Portfolio created successfully", "data": createPortfolioResponse})

}

func UpdatePortfolio(c *gin.Context) {
	var portfolio models.Portfolio
	if err := c.ShouldBindBodyWithJSON(&portfolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid data provided" + err.Error()})
		return
	}
	idParam := c.Param("id")
	portfolioId, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid portfolio id"})
		return
	}

	var existingPortfolio models.Portfolio
	if err := database.DB.First(&existingPortfolio, portfolioId).Error; err != nil {

		fmt.Println(err.Error())

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Portfolio not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Database error", "details": err.Error()})
		}
		return
	}
	// var updatedPortfolio models.Portfolio
	result := database.DB.Model(&models.Portfolio{}).Where("id=?", portfolioId).Updates(portfolio)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to update portfolio", "details": result.Error.Error()})
		return
	}
	database.DB.First(&existingPortfolio, portfolioId)

	updateResponse := dto.PortfolioResponseDto{
		ID:        existingPortfolio.ID,
		Name:      existingPortfolio.Name,
		UserID:    existingPortfolio.UserID,
		CreatedAt: existingPortfolio.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: existingPortfolio.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Portfolio updated successfully", "data": updateResponse})
}

func DeletePortfolio(c *gin.Context) {
	idParam := c.Param("id")
	portfolioId, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid portfolio id"})
		return
	}
	var existingPortfolio models.Portfolio
	if err := database.DB.First(&existingPortfolio, portfolioId).Error; err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Portfolio not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Database error", "details": err.Error()})
		}
		return
	}

	result := database.DB.Select("Assets").Delete(&existingPortfolio)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to delete portfolio", "details": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Portfolio deleted successfully"})
}

func GetPortfolio(c *gin.Context) {
	idParam := c.Param("id")
	portfolioId, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid portfolio id"})
		return
	}

	var portfolio models.Portfolio
	dbError := database.DB.Model(&models.Portfolio{}).Preload("Assets").First(&portfolio, portfolioId).Error

	if dbError != nil {
		log.Println(dbError.Error())
		if errors.Is(dbError, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Portfolio not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Database error", "details": dbError.Error()})
		}
		return
	}

	response := dto.GetPortfolioResponseDto{
		ID:        portfolio.ID,
		Name:      portfolio.Name,
		UserID:    portfolio.UserID,
		Assets:    portfolio.Assets,
		CreatedAt: portfolio.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: portfolio.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": response})
}
