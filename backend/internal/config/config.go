package config

import (
	"crypto/rand"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

var JwtSecret = make([]byte, 64)

func init() {
	rand.Read(JwtSecret) // Generate JwtSecret
}

// DB_PATH, HASHED_PASSWORD, APP_LISTEN

func Config(key string) string {
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Error("Error loading .env file")
	}
	return os.Getenv(key)
}
