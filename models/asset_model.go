package models

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	PortfolioID uint         `json:"portfolio_id"`
	Name        string       `json:"name"`
	Value       float64      `json:"value"`
	History     []AssetValue `gorm:"foreignKey:AssetID;constraint:OnDelete:CASCADE" json:"history"`
}
