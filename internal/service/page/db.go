package page

import (
	"below/internal/server/database"
	"below/internal/service"

	"gorm.io/gorm/clause"
)

func InitDB() error {
	return database.DB.
		AutoMigrate(&service.Page{}, &service.Site{})
}

func GetAllSiteDB() ([]*service.Site, error) {
	var sites []*service.Site
	err := database.DB.
		Find(&sites).Error
	if err != nil {
		return nil, err
	}
	return sites, nil
}

func QuerySiteDB(q *service.ReqSite) (*service.Site, error) {
	var site service.Site
	err := database.DB.
		Preload("Pages").
		Where("host = ?", q.Site).
		First(&site).Error
	return &site, err
}

func QueryPageDB(q *service.ReqPage) (*service.Page, error) {
	var page service.Page
	err := database.DB.
		Preload("Comments").
		Table("pages").
		Joins("JOIN sites ON pages.site_id = sites.id").
		Where("pages.path = ? AND sites.host = ?", q.Path, q.Site).
		First(&page).Error
	if err != nil {
		return nil, err
	}
	return &page, nil
}

func DelPageDB(q *service.ReqPage) error {
	page, err := QueryPageDB(q)
	if err != nil {
		return err
	}
	return database.DB.
		Unscoped().
		Select(clause.Associations).
		Delete(page).Error
}

func DelSiteDB(q *service.ReqSite) error {
	site, err := QuerySiteDB(q)
	if err != nil {
		return err
	}
	err = database.DB.
		Unscoped().
		Select(clause.Associations).
		Delete(site).Error
	return err
}

func InsertPagesDB(s *service.Site, p []service.Page) (int64, error) {
	res := database.DB.
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(s)
	if res.Error != nil {
		return 0, res.Error
	}

	// Get SiteID for pages we gonna add
	var siteID uint
	if res.RowsAffected == 0 {
		// if exist then find site_id via Host
		var site service.Site
		database.DB.
			Where(s, "Host").
			First(&site)
		siteID = site.ID
	} else {
		// if not then use the id we just added
		siteID = s.ID
	}

	for i := range p {
		p[i].SiteID = siteID
	}

	res = database.DB.
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(&p)

	return res.RowsAffected, res.Error
}
