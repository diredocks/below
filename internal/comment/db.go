package comment

import (
	"below/internal/database"
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

func QueryDB(q *CommentQueryByPage) ([]Comment, error) {
	res := []Comment{}
	err := database.Engine.Where("site = ? AND page = ? AND status = 'Sent'", q.Site, q.Page).Find(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteByIdDB(q *CommentQueryByID) (int64, error) {
	affected, err := database.Engine.In("id", q.IDs).Delete(&Comment{})
	if err != nil {
		return 0, err
	}
	return affected, nil
}

func DeleteByPageDB(q *CommentQueryByPage) (int64, error) {
	affected, err := database.Engine.Where("site = ? AND page = ? AND status = 'Sent'", q.Site, q.Page).Delete(&Comment{})
	if err != nil {
		return 0, err
	}
	return affected, nil
}
