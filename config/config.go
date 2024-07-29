package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port           = Config("PORT")
	DB_USER        = Config("DB_USER")
	DB_PASS        = Config("DB_PASS")
	DB_HOST        = Config("DB_HOST")
	DB_PORT        = Config("DB_PORT")
	DB_NAME        = Config("DB_NAME")
	ACCESS_SECRET  = Config("JWT_ACCESS_SECRET")
	REFRESH_SECRET = Config("JWT_REFRESH_SECRET")
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}
