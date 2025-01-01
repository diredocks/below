package comment

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func Validator[T any](model *T) fiber.Handler {
	return func(c *fiber.Ctx) error {
		instance := new(T)
		// Parse the request body into the model
		if err := c.BodyParser(instance); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Validate the model
		if err := validate.Struct(instance); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// Pass the validated model to the next handler via context
		c.Locals("validatedBody", instance)
		return c.Next()
	}
}

func Index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusTeapot).SendString("I'm not a Tea pot!")
}

func Add(c *fiber.Ctx) error {
	com := c.Locals("validatedBody").(*Comment)

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
	q := c.Locals("validatedBody").(*CommentQueryByPage)
	res, err := QueryDB(q)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch comments",
		})
	}

	return c.JSON(res)
}

func DelById(c *fiber.Ctx) error {
	q := c.Locals("validatedBody").(*CommentQueryByID)
	affected, err := DeleteByIdDB(q)
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

func DelByPage(c *fiber.Ctx) error {
	q := c.Locals("validatedBody").(*CommentQueryByPage)

	affected, err := DeleteByPageDB(q)
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
