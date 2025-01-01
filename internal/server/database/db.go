package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	// TODO: configurable db path
	DB, err = gorm.Open(sqlite.Open("./below.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
