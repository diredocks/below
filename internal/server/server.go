package server

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"below/internal/auth"
	"below/internal/comment"
	"below/internal/server/database"
)

func New(ctx context.Context) (*fiber.App, error) {
	// Initialize database
	dbInit := []func() error{
		database.InitDB, // This initialize db engine
		comment.InitDB,
		auth.InitDB,
	}
	for _, init := range dbInit {
		if err := init(); err != nil {
			return nil, err
		}
	}

	// Config middleware and misc
	app := fiber.New(fiber.Config{
		AppName: "Below",
	})
	app.Use(logger.New(logger.Config{
		TimeFormat: time.RFC3339,
		Format:     "[${time}] ${ip}:${port} ${status} - ${method} ${path}\n",
	}))

	// Config route
	api_router := app.Group("/api")
	comment.Router(api_router)
	auth.Router(api_router)
	return app, nil
}
