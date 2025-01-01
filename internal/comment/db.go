package comment

import (
	"below/internal/server/database"
)

func InitDB() error {
	// AutoMigrate syncs the schema with the database
	return database.DB.AutoMigrate(&Comment{})
}

func InsertDB(c *Comment) error {
	c.Status = StatusSent // Default to StatusSent
	return database.DB.Create(c).Error
}

func QueryDB(q *CommentQueryByPage) ([]Comment, error) {
	var res []Comment
	err := database.DB.Where("site = ? AND page = ? AND status = ?", q.Site, q.Page, StatusSent).Find(&res).Error
	return res, err
}

func DeleteByIdDB(q *CommentQueryByID) (int64, error) {
	result := database.DB.Delete(&Comment{}, q.IDs)
	return result.RowsAffected, result.Error
}

func DeleteByPageDB(q *CommentQueryByPage) (int64, error) {
	result := database.DB.Where("site = ? AND page = ? AND status = ?", q.Site, q.Page, StatusSent).Delete(&Comment{})
	return result.RowsAffected, result.Error
}
