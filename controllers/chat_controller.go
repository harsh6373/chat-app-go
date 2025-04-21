package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/harsh6373/chat-app-go/services"
)

func GetMessages(c *fiber.Ctx) error {
	userID, _ := strconv.Atoi(c.Params("userID"))
	peerID, _ := strconv.Atoi(c.Params("peerID"))

	messages, err := services.GetChatHistory(uint(userID), uint(peerID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(messages)
}
