package dto

import "github.com/AVtheking/user_portfolio_management/models"

type SignUpRequestDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type UserResponseDto struct {
	ID           uint   `json:"id"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	AccessToken  string `json:"access_token "`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequestDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUserResponseDto struct {
	ID         uint               `json:"id"`
	Email      string             `json:"email"`
	Username   string             `json:"username"`
	Portfolios []models.Portfolio `json:"portfolios"`
	CreatedAt  string             `json:"created_at"`
	UpdatedAt  string             `json:"updated_at"`
}

type CreatePortfolioRequestDto struct {
	Name string `json:"name" binding:"required"`
}

type PortfolioResponseDto struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetPortfolioResponseDto struct {
	ID        uint           `json:"id"`
	UserID    uint           `json:"user_id"`
	Name      string         `json:"name"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	Assets    []models.Asset `json:"assets"`
}

type CreateAssetRequestDto struct {
	Name  string  `json:"name" binding:"required"`
	Value float64 `json:"value" binding:"required"`
}

type AssetResponseDto struct {
	ID          uint    `json:"id"`
	PortfolioID uint    `json:"portfolio_id"`
	Name        string  `json:"name"`
	Value       float64 `json:"value"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type AssetHistory struct {
	Value float64 `json:"value"`
	Date  string  `json:"date"`
}

type GetAssetResponseDto struct {
	ID          uint           `json:"id"`
	PortfolioID uint           `json:"portfolio_id"`
	Name        string         `json:"name"`
	Value       float64        `json:"value"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
	History     []AssetHistory `json:"asset_history"`
}
