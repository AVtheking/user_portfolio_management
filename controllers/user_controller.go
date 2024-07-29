package controller

import (
	"log"
	"net/http"

	"github.com/AVtheking/user_portfolio_management/database"
	"github.com/AVtheking/user_portfolio_management/dto"
	"github.com/AVtheking/user_portfolio_management/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, exists := c.Get("userId")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User not authenticated"})

	}

	var user models.User
	result := database.DB.Model(&models.User{}).Preload("Portfolios").First(&user, "id = ?", userId)
	if result.Error != nil {
		log.Println(result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": result.Error.Error()})
		return
	}

	userResponse := dto.GetUserResponseDto{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		Portfolios: user.Portfolios,
		CreatedAt:  user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": userResponse})
}
