package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusTeapot).SendString("I'm not a Tea pot!")
}

func Login(c *fiber.Ctx) error {
	return c.SendString("Login")
}

func Logout(c *fiber.Ctx) error {
	return c.SendString("Cancel")
}

func Register(c *fiber.Ctx) error {
	return c.SendString("Register")
}

func Cancel(c *fiber.Ctx) error {
	return c.SendString("Canle")
}
