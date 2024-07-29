package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AVtheking/user_portfolio_management/database"
	"github.com/AVtheking/user_portfolio_management/dto"
	"github.com/AVtheking/user_portfolio_management/models"
	"github.com/AVtheking/user_portfolio_management/utils"
	"github.com/gin-gonic/gin"
)


func SignUp(c *gin.Context) {
	var request dto.SignUpRequestDto

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request"})
		// Log the error internally for further investigation
		log.Printf("Error binding JSON: %v", err)
		return
	}
	var user models.User

	result := database.DB.Find(&user, "email = ?", request.Email)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": result.Error.Error()})
		return
	}

	if user.Email != "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "User with the email already exists"})
		return
	}

	hashPassword := utils.HashPassword(request.Password)

	newUser := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashPassword,
	}

	createResult := database.DB.Create(&newUser)

	if createResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "An error occurred creating user"})
		return
	}
	access_token, err := utils.GenerateAccessToken(newUser.ID, newUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err})
		return
	}
	refresh_token, err := utils.GenerateRefreshToken(newUser.ID, newUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err})
		return
	}
	fmt.Printf("User created successfully %v", newUser)
	userResponse := dto.UserResponseDto{
		ID:           newUser.ID,
		Email:        newUser.Email,
		Username:     newUser.Username,
		CreatedAt:    newUser.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    newUser.UpdatedAt.Format("2006-01-02 15:04:05"),
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "User Registered successfully", "data": userResponse})

}

func Login(c *gin.Context) {
	var request dto.LoginRequestDto

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid data provided"})
		return
	}
	var user models.User
	result := database.DB.Find(&user, "email = ?", request.Email)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": result.Error.Error()})
		return
	}

	if user.Email == "" {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "User not found"})
		return
	}
	passwordMatched := utils.ComparePassword(user.Password, request.Password)

	if !passwordMatched {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid login credentials"})
		return
	}

	accessToken, err := utils.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "An error occurred generating access token"})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "An error occurred generating refresh token"})
		return
	}
	userResponse := dto.UserResponseDto{
		ID:           user.ID,
		Email:        user.Email,
		Username:     user.Username,
		CreatedAt:    user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    user.UpdatedAt.Format("2006-01-02 15:04:05"),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "User logged in successfully", "data": userResponse})
}
