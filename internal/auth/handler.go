package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusTeapot).SendString("I'm not a Tea pot!")
}
