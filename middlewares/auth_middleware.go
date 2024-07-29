package middlewares

import (
	"net/http"
	"strings"

	"github.com/AVtheking/user_portfolio_management/config"
	"github.com/AVtheking/user_portfolio_management/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleWare function to check if the user is authenticated
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Authorization header required"})
			c.Abort()
			return
		}
		tokenString := ""
		if strings.Contains(authHeader, "Bearer") {
			tokenString = strings.TrimSpace(strings.Split(authHeader, "Bearer")[1])
		}
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Token required"})
			c.Abort()
			return
		}
		claims := &utils.Claim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.ACCESS_SECRET), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("userId", claims.ID)
		c.Set("userEmail", claims.Email)
		c.Next()

	}

}
