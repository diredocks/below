package comment

import (
	"time"
)

type CommentStatus string

const (
	StatusSent    CommentStatus = "Sent"
	StatusRemoved CommentStatus = "Removed"
)

type Comment struct {
	ID        uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	Site      string        `gorm:"type:text;not null" validate:"required,url" json:"site"`
	Page      string        `gorm:"type:text;not null" validate:"required,uri" json:"page"`
	Nickname  string        `gorm:"type:text;not null" validate:"required" json:"nickname"`
	Email     string        `gorm:"type:text" validate:"omitempty,email" json:"email"`
	Content   string        `gorm:"type:text;not null" validate:"required" json:"content"`
	CreatedAt time.Time     `gorm:"autoCreateTime" json:"created_at"`
	Status    CommentStatus `gorm:"type:text;not null;default:'Sent'" json:"status"`
}

type CommentQueryByPage struct {
	Site string `validate:"required,url" json:"site"`
	Page string `validate:"required,uri" json:"page"`
}

type CommentQueryByID struct {
	IDs []string `validate:"gt=0,dive,number,required" json:"ids"`
}
