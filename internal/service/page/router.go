package page

import (
	"below/internal/server/middleware"
	"below/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Router(app fiber.Router) {
	group := app.Group("/page")
	group.Post("/", Index)
	group.Post("/add", middleware.Validator(&service.Page{}), Add)
	group.Post("/get", middleware.Validator(&service.Page{}), Get)
	group.Post("/del", middleware.Validator(&service.Page{}), Del)
}
