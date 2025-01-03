package config

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

// DB_PATH, HASHED_PASSWORD, APP_LISTEN

func Config(key string) string {
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Error("Error loading .env file")
	}
	return os.Getenv(key)
}
