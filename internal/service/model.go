package service

import (
	"gorm.io/gorm"
)

type CommentStatus string

const (
	StatusSent    CommentStatus = "Sent"
	StatusPending CommentStatus = "Pending"
)

type Site struct {
	gorm.Model
	Host    string `gorm:"type:text;not null;size=128;uniqueIndex" validate:"required,url,max=128" json:"site"`
	SiteMap string `gorm:"type:text;not null;size=128" validate:"required,max=128" json:"site_map"`
	Pages   []Page `gorm:"constraint:onDelete:CASCADE"`
}

type Page struct {
	gorm.Model
	Path string `gorm:"type:text;not null;size=255;uniqueIndex:idx_site_path" validate:"required,uri,max=255" json:"path"`
	// TODO: views of page
	Comments []Comment `gorm:"constraint:onDelete:CASCADE"`
	SiteID   uint      `gorm:"uniqueIndex:idx_site_path" validate:"required,number" json:"site_id"`
}

type Comment struct {
	gorm.Model
	Name    string `gorm:"type:text;not null;size=64" validate:"required,min=3,max=64" json:"name"`
	Content string `gorm:"type:text;not null;size=1024" validate:"required,max=1024" json:"content"`
	Email   string `gorm:"type:text;not null;size=64" validate:"omitempty,email,max=64" json:"email"`
	// TODO: upvote and downvote of comment
	Status CommentStatus `gorm:"type:text;not null;default:'Sent'"`
	PageID uint          `validate:"required,number" json:"page_id"`
}

type ReqSiteMap struct {
	SiteMap string `validate:"required,max=128" json:"site_map"`
}

type ReqSite struct {
	Site string `validate:"required,max=128" json:"site"`
}

type ReqPage struct {
	ReqSite
	Path string `validate:"required,uri,max=255" json:"path"`
}

type ReqID struct {
	ID uint `validate:"gt=0,number,required" json:"id"`
}

type ReqIDs struct {
	IDs []uint `validate:"gt=0,dive,number,required" json:"ids"`
}
