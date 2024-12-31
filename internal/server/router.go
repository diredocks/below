package server

import (
	"below/internal/comment"

	"github.com/gofiber/fiber/v2"
)

func CommentApiRouter(app fiber.Router) {
	router_group := app.Group("/api")
	router_group.Get("/", comment.IndexHandler)
	router_group.Post("/add", comment.AddHandler)
	router_group.Post("/get", comment.GetHandler)
}
