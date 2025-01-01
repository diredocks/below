package middleware

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
