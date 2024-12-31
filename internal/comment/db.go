package comment

import (
	"below/internal/database"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() error {
	// Sync database with struct Comment
	if err := database.Engine.Sync2(new(Comment)); err != nil {
		return err
	}
	return nil
}

func InsertDB(c *Comment) error {
	c.Status = StatusSent // Default to StatusSent
	_, err := database.Engine.Insert(c)
	return err
}

func QueryDB(q *CommentQuery) ([]Comment, error) {
	res := []Comment{}
	err := database.Engine.Where("site = ? AND page = ?", q.Site, q.Page).Find(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
