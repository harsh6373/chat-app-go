package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/harsh6373/chat-app-go/utils"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(utils.JWTSecret()),
		ContextKey:   "user",
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}
