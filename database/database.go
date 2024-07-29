package database

import (
	"fmt"
	"strconv"

	"github.com/AVtheking/user_portfolio_management/config"
	"github.com/AVtheking/user_portfolio_management/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	port, err := strconv.ParseInt(config.DB_PORT, 10, 32)

	if err != nil {
		panic(err)
	}

	configData := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", config.DB_HOST, port, config.DB_USER, config.DB_NAME, config.DB_PASS)

	var dbErr error
	DB, dbErr = gorm.Open(postgres.Open(configData), &gorm.Config{})
	if dbErr != nil {
		panic("Error connecting to database: " + dbErr.Error())
	}

	fmt.Println("\x1b[32m...............Database connected..................\x1b[0m")

	DB.AutoMigrate(&models.User{}, &models.Portfolio{}, &models.Asset{}, &models.AssetValue{})
}
