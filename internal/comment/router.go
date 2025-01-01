package comment

import (
	"below/internal/server/middleware"

	"github.com/gofiber/fiber/v2"
)

func Router(app fiber.Router) {
	router_group := app.Group("/comment")
	router_group.Get("/", Index)
	router_group.Post("/get", middleware.Validator(&CommentQueryByPage{}), Get)
	router_group.Post("/add", middleware.Validator(&Comment{}), Add)
	router_group.Post("/delete/id", middleware.Validator(&CommentQueryByID{}), DelById)
	router_group.Post("/delete/page", middleware.Validator(&CommentQueryByPage{}), DelByPage)
}
