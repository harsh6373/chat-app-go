// ws/broadcast_test.go
package ws

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	fiberWs "github.com/gofiber/websocket/v2"
)

// MockWSHandler is a simplified version of HandleWebSocket for testing
// that doesn't rely on database operations
func MockWSHandler(c *fiberWs.Conn) {
	userID := c.Params("userID")
	client := &Client{Conn: c, ID: userID}

	mu.Lock()
	clients[userID] = client
	mu.Unlock()

	fmt.Println("User connected:", userID)

	for {
		var msg IncomingMessage
		if err := c.ReadJSON(&msg); err != nil {
			fmt.Println("Error reading JSON:", err)
			break
		}

		// Skip database operations for testing

		// Handle broadcast
		if msg.Broadcast {
			broadcast(userID, msg.Text)
		} else if msg.Receiver != "" {
			sendToClient(msg.Receiver, fmt.Sprintf("From %s: %s", userID, msg.Text))
		}
	}

	mu.Lock()
	delete(clients, userID)
	mu.Unlock()
	fmt.Println("User disconnected:", userID)
}

func TestBroadcast(t *testing.T) {
	// Clear any existing clients
	mu.Lock()
	for k := range clients {
		delete(clients, k)
	}
	mu.Unlock()

	// Setup Fiber app with WebSocket route
	app := fiber.New()
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol
		if fiberWs.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/ws/:userID", fiberWs.New(MockWSHandler)) // Use mock handler

	// Start the server in a goroutine
	go func() {
		err := app.Listen(":3100")
		if err != nil {
			// Only log if it's not a server closed error (which happens on shutdown)
			if !strings.Contains(err.Error(), "server closed") {
				t.Logf("Server error: %v", err)
			}
		}
	}()
	// Give the server time to start
	time.Sleep(100 * time.Millisecond)
	defer app.Shutdown()

	// Create WebSocket URL
	url := "ws://localhost:3100/ws/"

	// Connect two clients using the standard dialer
	dialer := &websocket.Dialer{}
	c1, _, err := dialer.Dial(url+"user1", nil)
	if err != nil {
		t.Fatalf("Failed to connect client1: %v", err)
	}
	defer c1.Close()

	c2, _, err := dialer.Dial(url+"user2", nil)
	if err != nil {
		t.Fatalf("Failed to connect client2: %v", err)
	}
	defer c2.Close()

	// Give the connections time to establish
	time.Sleep(100 * time.Millisecond)

	// Send message from client1
	message := map[string]interface{}{"text": "Hello from user1", "broadcast": true}
	if err := c1.WriteJSON(message); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Read message on client2
	var received string
	if err := c2.ReadJSON(&received); err != nil {
		t.Fatalf("Failed to read message: %v", err)
	}

	// Check if the message contains the expected text
	if !strings.Contains(received, "Hello from user1") {
		t.Errorf("Expected message to contain 'Hello from user1', got '%s'", received)
	}
}
