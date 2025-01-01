package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Router(app fiber.Router) {
	router_group := app.Group("/auth")
	router_group.Get("/", Index)
	router_group.Post("/login", Login)
	router_group.Post("/logout", Logout)
	router_group.Post("/register", Register)
	router_group.Post("/cancel", Cancel)
}
