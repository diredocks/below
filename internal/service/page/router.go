package page

import (
	"below/internal/server/middleware"
	"below/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Router(app fiber.Router) {
	group := app.Group("/page")
	group.Get("/", Index)
	group.Post("/add", middleware.Validator(&service.Page{}), Add)
}
