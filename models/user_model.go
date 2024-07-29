package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string      `json:"username"`
	Password   string      `json:"password"`
	Email      string      `json:"email"`
	Portfolios []Portfolio `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"portfolios"`
}
