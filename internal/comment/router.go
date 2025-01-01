package comment

import (
	"github.com/gofiber/fiber/v2"
)

func Router(app fiber.Router) {
	router_group := app.Group("/comment")
	router_group.Get("/", Index)
	router_group.Post("/get", Validator(&CommentQueryByPage{}), Get)
	router_group.Post("/add", Validator(&Comment{}), Add)
	router_group.Post("/delete/id", Validator(&CommentQueryByID{}), DelById)
	router_group.Post("/delete/page", Validator(&CommentQueryByPage{}), DelByPage)
}
