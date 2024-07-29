package utils

import (
	"time"

	"github.com/AVtheking/user_portfolio_management/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claim struct {
	ID    uint
	Email string
	jwt.RegisteredClaims
}

// HashPassword hashes the password
func HashPassword(password string) string {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic("Failed to hash password:" + err.Error())
	}
	return string(hashPassword)
}

// ComparePassword compares the hashed password and the password
func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

// GenerateAccessToken generates an access token
func GenerateAccessToken(id uint, email string) (string, error) {
	expiryTime := time.Now().Add(1 * time.Hour)
	claims := Claim{
		ID:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiryTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(config.ACCESS_SECRET))

	return accessToken, err
}

// GenerateRefreshToken generates a refresh token
func GenerateRefreshToken(id uint, email string) (string, error) {
	expiryTime := time.Now().Add(7 * 24 * time.Hour)
	claims := Claim{
		ID:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiryTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshToken, err := token.SignedString([]byte(config.REFRESH_SECRET))

	return refreshToken, err
}
