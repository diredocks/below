package comment

import (
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func InitDB() error {
	var err error
	Engine, err = xorm.NewEngine("sqlite3", "./comments.db") // TODO: db path config
	if err != nil {
		return err
	}

	if err = Engine.Sync2(new(Comment)); err != nil {
		return err
	}

	return nil
}

func CommentInsertDB(com *Comment) error {
	com.Status = StatusSent
	_, err := Engine.Insert(com)
	return err
}

func CommentQueryDB(site, page string) ([]Comment, error) {
	var comments []Comment
	err := Engine.Where("site = ? AND page = ?", site, page).Find(&comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
