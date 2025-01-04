package page

import (
	"below/internal/server/middleware"
	"below/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Router(app fiber.Router) {
	page := app.Group("/page")
	page.Post("/", Index)
	// Get a page and all its comment by site and path
	page.Post("/get", middleware.Validator(&service.ReqPage{}), GetPage)
	// Delete a Page
	page.Post("/del", middleware.Protected(), middleware.Validator(&service.ReqPage{}), DelPage)

	site := app.Group("/site")
	// Update site
	site.Post("/update", middleware.Validator(&service.ReqSite{}), UpdateSite)
	site.Use(middleware.Protected())
	// Get all sites
	site.Post("/", GetAllSite)
	// Get a site and all its pages
	site.Post("/get", middleware.Validator(&service.ReqSite{}), GetSite)
	// Delete whole site
	site.Post("/del", middleware.Validator(&service.ReqSite{}), DelSite)
	// Add a new site
	site.Post("/add", middleware.Validator(&service.ReqSiteMap{}), AddSite)
}
