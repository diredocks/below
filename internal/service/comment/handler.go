package comment

import (
	"below/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusTeapot).SendString("I'm not a Tea pot!")
}

func Add(c *fiber.Ctx) error {
	com := c.Locals("validatedBody").(*service.Comment)

	if err := InsertDB(com); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to send comment",
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "comment sent",
	})
}
