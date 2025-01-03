package comment

import (
	"below/internal/server/middleware"
	"below/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Router(app fiber.Router) {
	group := app.Group("/comment")
	group.Get("/", Index)
	group.Post("/add", middleware.Validator(&service.Comment{}), Add)
	group.Post("/del", middleware.Protected(), middleware.Validator(&service.ReqIDs{}), Del)
}
