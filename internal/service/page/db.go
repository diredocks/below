package page

import (
	"below/internal/server/database"
	"below/internal/service"
)

func InitDB() error {
	return database.DB.AutoMigrate(&service.Page{})
}

func InsertDB(p *service.Page) error {
	return database.DB.Create(&p).Error
}

func QueryDB(p *service.Page) ([]service.Page, error) {
	var pages []service.Page
	err := database.DB.Model(&service.Page{}).Preload("Comments").Find(&pages).Error
	return pages, err
}
