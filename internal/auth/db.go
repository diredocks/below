package auth

import (
	_ "below/internal/server/database"
)

func InitDB() error {
	/*if err := database.Engine.Sync2(new(User)); err != nil {
		return err
	}*/
	return nil
}
