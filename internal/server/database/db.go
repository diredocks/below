package database

import (
	"below/internal/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.Config("DB_PATH")), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
