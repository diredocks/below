package page

import (
	"below/internal/server/middleware"
	"below/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Router(app fiber.Router) {
	page := app.Group("/page")
	page.Post("/", Index)
	page.Post("/get", middleware.Validator(&service.ReqPage{}), GetPage)        // Get a page and all its comment by site and path
	page.Post("/del", middleware.Validator(&service.ReqPage{}), DelPage)        // Delete a Page
	page.Post("/add", middleware.Validator(&service.ReqSiteMap{}), AddOrUpdate) // Update or Add a new site
	page.Post("/update", middleware.Validator(&service.ReqSiteMap{}), AddOrUpdate)
	site := app.Group("/site")
	site.Post("/", GetAllSite)
	site.Post("/get", middleware.Validator(&service.ReqSite{}), GetSite) // Get a site and all its pages
	site.Post("/del", middleware.Validator(&service.ReqSite{}), DelSite) // Delete whole site
}
