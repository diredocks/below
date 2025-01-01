package service

import (
	"gorm.io/gorm"
)

type CommentStatus string

const (
	StatusSent    CommentStatus = "Sent"
	StatusPending CommentStatus = "Pending"
)

type Page struct {
	gorm.Model
	Site     string    `gorm:"type:text;not null;size=128;uniqueIndex:idx_site_path" validate:"required,url,max=128" json:"site"`
	Path     string    `gorm:"type:text;not null;size=255;uniqueIndex:idx_site_path" validate:"required,uri,max=255" json:"path"`
	Comments []Comment `gorm:"constraint:onDelete:CASCADE"`
}

type Comment struct {
	gorm.Model
	Nickname string        `gorm:"type:text;not null;size=64" validate:"required,min=3,max=64" json:"nickname"`
	Content  string        `gorm:"type:text;not null;size=1024" validate:"required,max=1024" json:"content"`
	Email    string        `gorm:"type:text;not null;size=64" validate:"omitempty,email,max=64" json:"email"`
	Status   CommentStatus `gorm:"type:text;not null;default:'Sent'"`
	PageID   uint          `validate:"required,number" json:"page_id"`
}

type QueryByID struct {
	ID uint `validate:"gt=0,number,required" json:"id"`
}

type QueryByIDs struct {
	IDs []uint `validate:"gt=0,dive,number,required" json:"ids"`
}
