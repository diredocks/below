package page

import (
	"below/internal/server/database"
	"below/internal/service"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func InitDB() error {
	return database.DB.AutoMigrate(&service.Page{})
}

func InsertDB(p *service.Page) error {
	return database.DB.Create(&p).Error
}

func QueryDB(p *service.Page) (service.Page, error) {
	var page service.Page
	err := database.DB.
		Where(p, "Site", "Path").
		Preload("Comments").
		First(&page).Error
	return page, err
}

func GetAllDB(p *service.Page) ([]service.Page, error) {
	var pages []service.Page
	err := database.DB.Preload("Comments").Find(&pages).Error
	return pages, err
}

func DelDB(p *service.Page) error {
	query, err := QueryDB(p)
	if err != nil {
		return err
	}
	res := database.DB.Unscoped().
		Select(clause.Associations).Delete(&service.Page{
		Model: gorm.Model{
			ID: query.ID,
		},
	})
	return res.Error
}
