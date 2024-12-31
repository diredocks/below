package server

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func NewServer(ctx context.Context) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Below",
	})

	CommentApiRouter(app)
	return app
}
