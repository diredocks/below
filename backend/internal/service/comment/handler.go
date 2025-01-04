package comment

import (
	"below/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusTeapot).
		SendString("I'm not a Tea pot!")
}

func Add(c *fiber.Ctx) error {
	com := c.Locals("validatedBody").(*service.Comment)
	com.Status = service.StatusPending // Defalut to Pending

	if err := InsertDB(com); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "failed to send comment",
				"msg":   err.Error(),
			})
	}

	return c.Status(fiber.StatusCreated).
		JSON(fiber.Map{
			"success": "comment sent",
		})
}

func Del(c *fiber.Ctx) error {
	q := c.Locals("validatedBody").(*service.ReqIDs)
	affected, err := DelDB(q)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "failed to delete comments",
				"msg":   err.Error(),
			})
	}

	return c.JSON(fiber.Map{
		"success":  "deleted comment(s)",
		"affected": affected,
	})
}
