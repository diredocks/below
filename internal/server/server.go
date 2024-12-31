package server

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"below/internal/comment"
	"below/internal/database"
)

func New(ctx context.Context) (*fiber.App, error) {
	// Initialize database
	database.InitDB()
	comment.InitDB()

	// Config middleware and misc
	app := fiber.New(fiber.Config{
		AppName: "Below",
	})
	app.Use(logger.New(logger.Config{
		TimeFormat: time.RFC3339,
		Format:     "[${time}] ${ip}:${port} ${status} - ${method} ${path}\n",
	}))

	// Config route
	CommentApiRouter(app)
	return app, nil
}
