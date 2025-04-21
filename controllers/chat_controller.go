// controllers/chat_controller.go
package controllers

import (
    "github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"message": "login logic here"})
}

func Register(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"message": "register logic here"})
}
