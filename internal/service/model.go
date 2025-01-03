package service

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentStatus string

const (
	StatusApproved CommentStatus = "Approved"
	StatusPending  CommentStatus = "Pending"
	StatusHided    CommentStatus = "Hided"
)

type Site struct {
	gorm.Model
	Host    string `gorm:"type:text;not null;size=128;uniqueIndex" validate:"required,url,max=128" json:"site"`
	SiteMap string `gorm:"type:text;not null;size=128" validate:"required,max=128" json:"site_map"`
	Pages   []Page `gorm:"constraint:OnDelete:CASCADE"`
}

type Page struct {
	gorm.Model
	Path string `gorm:"type:text;not null;size=255;uniqueIndex:idx_site_path" validate:"required,uri,max=255" json:"path"`
	// TODO: Views    uint      `gorm:"default:0"`
	Comments []Comment `gorm:"constraint:OnDelete:CASCADE"`
	SiteID   uint      `gorm:"uniqueIndex:idx_site_path" validate:"required,number" json:"site_id"`
}

type Comment struct {
	gorm.Model
	Name    string `gorm:"type:text;not null;size=64" validate:"required,min=3,max=64" json:"name"`
	Content string `gorm:"type:text;not null;size=1024" validate:"required,max=1024" json:"content"`
	Email   string `gorm:"type:text;not null;size=64" validate:"omitempty,email,max=64" json:"email"`
	// TODO: Upvotes   uint          `gorm:"default:0"`
	// TODO: Downvotes uint          `gorm:"default:0"`
	Status CommentStatus `gorm:"type:text;not null;default:'Pending'"`
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

func (s *Site) BeforeDelete(tx *gorm.DB) (err error) {
	var pages = []Page{}
	if err := tx.
		Where("site_id = ?", s.ID).
		Find(&pages).Error; err != nil {
		return err
	}
	if err := tx.
		Unscoped().
		Select(clause.Associations).
		Delete(&pages).Error; err != nil {
		return err
	}
	return nil
}
