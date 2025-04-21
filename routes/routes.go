package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/harsh6373/chat-app-go/controllers"
	"github.com/harsh6373/chat-app-go/middleware"
	"github.com/harsh6373/chat-app-go/ws"
)

func SetupRoutes(app *fiber.App) {
	// Public endpoints
	app.Post("/login", controllers.Login)
	app.Post("/register", controllers.Register)

	// Protected WebSocket route
	app.Use("/ws/:userID", middleware.Protected()) // Ensures JWT is verified

	// WebSocket handler for real-time chat
	app.Get("/ws/:userID", websocket.New(ws.HandleWebSocket))
}
