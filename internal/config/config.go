package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
