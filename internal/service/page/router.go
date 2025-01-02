package page

import (
	"below/internal/server/middleware"
	"below/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Router(app fiber.Router) {
	page := app.Group("/page")
	page.Post("/", Index)
	page.Post("/del", middleware.Validator(&service.ReqPage{}), DelPage)
	page.Post("/del/all", middleware.Validator(&service.ReqSite{}), DelSite)
	page.Post("/get", middleware.Validator(&service.ReqPage{}), GetPage)        // Get a page and all its comment by site and path
	page.Post("/add", middleware.Validator(&service.ReqSiteMap{}), AddOrUpdate) // Update or Add a new site
	page.Post("/update", middleware.Validator(&service.ReqSiteMap{}), AddOrUpdate)
}
