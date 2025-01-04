package auth

import (
	"below/internal/server/middleware"
	"below/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Router(app fiber.Router) {
	auth := app.Group("/auth")
	auth.Post("/login", middleware.Validator(&service.ReqLogin{}), Login)
}
