package server

import (
	"below/internal/comment"

	"github.com/gofiber/fiber/v2"
)

func CommentApiRouter(app fiber.Router) {
	router_group := app.Group("/api")
	router_group.Get("/", comment.Index)
	router_group.Post("/comment/get", comment.Get)
	router_group.Post("/comment/add", comment.Add)
	router_group.Post("/comment/delete", comment.Del)
}
