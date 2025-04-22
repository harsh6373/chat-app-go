package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/harsh6373/chat-app-go/config"
	"github.com/harsh6373/chat-app-go/routes"
	"github.com/harsh6373/chat-app-go/ws"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.ConnectRedis()

	go ws.SubscribeTypingUpdates()

	app := fiber.New()

	// WebSocket middleware setup
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// WebSocket route
	app.Get("/ws/:userID", websocket.New(ws.HandleWebSocket))

	// API routes
	routes.SetupRoutes(app)

	app.Listen(":" + config.GetEnv("PORT", "3000"))
}
