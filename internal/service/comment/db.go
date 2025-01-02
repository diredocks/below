package comment

import (
	"below/internal/server/database"
	"below/internal/service"
)

func InitDB() error {
	return database.DB.AutoMigrate(&service.Comment{})
}

func InsertDB(c *service.Comment) error {
	if err := database.DB.First(&service.Page{}, c.PageID).Error; err != nil {
		return err
	}
	c.Status = service.StatusSent // Default to StatusSent
	return database.DB.Create(&c).Error
}

func DelDB(q *service.ReqIDs) (int64, error) {
	res := database.DB.Unscoped().Delete(&service.Comment{}, q.IDs)
	return res.RowsAffected, res.Error
}
