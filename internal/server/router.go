package server

import (
	"github.com/gofiber/fiber/v2"
)

func CommentApiRouter(app fiber.Router) {
	router_group := app.Group("/api")
	router_group.Get("/", CommentIndexHandler)
	router_group.Post("/add", CommentAddHandler)
	router_group.Post("/get", CommentGetHandler)
}
