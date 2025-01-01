package database

import (
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

// TODO: make database path configurable
func InitDB() error {
	var err error
	if Engine, err = xorm.NewEngine("sqlite3", "./below.db"); err != nil {
		return err
	}
	return nil
}
