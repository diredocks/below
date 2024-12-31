package comment

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func Index(c *fiber.Ctx) error {
	return c.SendString("Hello, Below!")
}

func Add(c *fiber.Ctx) error {
	com := new(Comment)
	if err := c.BodyParser(com); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	if err := validate.Struct(com); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "validation failed",
			"msg":   err.Error(),
		})
	}

	if err := InsertDB(com); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to save comment",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "comment sent",
	})
}

func Get(c *fiber.Ctx) error {
	q := new(CommentQueryByPage)
	if err := c.BodyParser(q); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid query"})
	}

	if err := validate.Struct(q); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "validation failed",
			"msg":   err.Error(),
		})
	}

	res, err := QueryDB(q)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch comments",
		})
	}

	return c.JSON(res)
}

func Del(c *fiber.Ctx) error {
	q := new(CommentQueryByID)
	if err := c.BodyParser(q); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid query"})
	}

	if err := validate.Struct(q); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "validation failed",
			"msg":   err.Error(),
		})
	}

	affected, err := DeleteDB(q)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete comments",
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success":  "deleted comment(s)",
		"affected": affected,
	})
}
