package page

import (
	"below/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusTeapot).SendString("I'm not a Tea pot!")
}

func Add(c *fiber.Ctx) error {
	p := c.Locals("validatedBody").(*service.Page)

	if err := InsertDB(p); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to add page",
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "page added",
	})
}
