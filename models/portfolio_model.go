package models

import (
	"gorm.io/gorm"
)

type Portfolio struct {
	gorm.Model
	UserID uint    `json:"user_id"`
	Name   string  `json:"name"`
	Assets []Asset `gorm:"foreignKey:PortfolioID;constraint:OnDelete:CASCADE" json:"assets"`
}
