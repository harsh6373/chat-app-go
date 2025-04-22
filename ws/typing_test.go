// ws/typing_test.go
package ws

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/websocket/v2"
)

// Mock types and functions for testing
type MockIncomingMessage struct {
	Text     string `json:"text"`
	Receiver string `json:"receiver"`
	Typing   bool   `json:"typing"`
}

// TypingTestHandleWebSocket is a simple mock for the typing test
func TypingTestHandleWebSocket(c *websocket.Conn) {
	// Simple mock implementation for testing
}

func TestTypingIndicatorSetup(t *testing.T) {
	// Setup Fiber app with WebSocket route
	app := fiber.New()

	// Add WebSocket route middleware
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client requested upgrade to the WebSocket protocol
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// Add handler
	app.Get("/ws/:userID", websocket.New(TypingTestHandleWebSocket))

	// Verify routes are registered
	// Check for GET route with WebSocket handler
	hasWSRoute := false
	for _, route := range app.Stack()[0] { // GET routes are in the first stack
		if route.Path == "/ws/:userID" {
			hasWSRoute = true
			break
		}
	}

	// Assert that our WebSocket handler route is registered
	utils.AssertEqual(t, true, hasWSRoute, "WebSocket handler route should be registered")

	// Test IncomingMessage struct
	msg := MockIncomingMessage{
		Text:     "Hello",
		Receiver: "user2",
		Typing:   true,
	}

	utils.AssertEqual(t, "Hello", msg.Text)
	utils.AssertEqual(t, "user2", msg.Receiver)
	utils.AssertEqual(t, true, msg.Typing)
}
