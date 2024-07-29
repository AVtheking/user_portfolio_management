package models

import "gorm.io/gorm"

type 
AssetValue struct {
	gorm.Model
	AssetID uint    `json:"asset_id"`
	Value   float64 `json:"value"`
}
