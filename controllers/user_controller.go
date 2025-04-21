package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harsh6373/chat-app-go/services"
	"github.com/harsh6373/chat-app-go/utils"
)

type AuthInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	var input AuthInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	user, err := services.CreateUser(input.Username, input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, _ := utils.GenerateJWT(user.ID)

	return c.JSON(fiber.Map{
		"token": token,
		"user":  user,
	})
}

func Login(c *fiber.Ctx) error {
	var input AuthInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	user, err := services.AuthenticateUser(input.Username, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	token, _ := utils.GenerateJWT(user.ID)

	return c.JSON(fiber.Map{
		"token": token,
		"user":  user,
	})
}
