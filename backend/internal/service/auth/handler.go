package auth

import (
	"below/internal/config"
	"below/internal/service"
	"time"

	"crypto/md5"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {
	q := c.Locals("validatedBody").(*service.ReqLogin)
	hashed_password := md5.Sum([]byte(q.Password))
	hash_string := hex.EncodeToString(hashed_password[:])
	if hash_string != config.Config("HASHED_PASSWORD") {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 2).Unix(), // expired in 2 hours
	})
	t, err := token.SignedString(config.JwtSecret)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"token": t})
}
