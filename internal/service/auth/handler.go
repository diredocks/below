package auth

import (
	"below/internal/config"
	"below/internal/service"

	"crypto/md5"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	q := c.Locals("validatedBody").(*service.ReqLogin)
	hashed_password := md5.Sum([]byte(q.Password))
	hash_string := hex.EncodeToString(hashed_password[:])
	if hash_string == config.Config("HASHED_PASSWORD") {
		return c.SendString("good!")
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}
