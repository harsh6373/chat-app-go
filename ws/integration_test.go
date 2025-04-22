// ws/integration_test.go
package ws

import (
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	fiberWs "github.com/gofiber/websocket/v2"
)

// MockHandler is a simplified version of HandleWebSocket for testing
func MockHandleWebSocket(c *fiberWs.Conn) {
	// Simple echo logic for testing
	for {
		var msg map[string]string
		if err := c.ReadJSON(&msg); err != nil {
			break
		}
		if err := c.WriteJSON(msg); err != nil {
			break
		}
	}
}

// TestWebSocketIntegration tests basic WebSocket functionality
func TestWebSocketIntegration(t *testing.T) {
	// Create a fiber app
	app := fiber.New()

	// Add WebSocket middleware for proper upgrades
	app.Use("/ws", func(c *fiber.Ctx) error {
		if fiberWs.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// Register our mock handler
	app.Get("/ws/:userID", fiberWs.New(MockHandleWebSocket))

	// Create a test server with a random port
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}

	// Start the server in a goroutine
	go func() {
		if err := app.Listener(ln); err != nil && err != http.ErrServerClosed {
			t.Logf("Server error: %v", err)
		}
	}()

	// Cleanup when test completes
	defer app.Shutdown()

	// Get the assigned port
	addr := ln.Addr().String()
	if !strings.Contains(addr, ":") {
		t.Fatalf("Invalid address: %s", addr)
	}

	// Create WebSocket URL
	wsURL := "ws://" + addr + "/ws/testuser"

	// Allow server to start
	time.Sleep(100 * time.Millisecond)

	// Connect to the WebSocket server
	dialer := &websocket.Dialer{}
	conn, _, err := dialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	defer conn.Close()

	// Send a test message
	testMsg := map[string]string{"text": "Hello WebSocket"}
	if err := conn.WriteJSON(testMsg); err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Read the echo response
	var response map[string]string
	if err := conn.ReadJSON(&response); err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	// Verify the response
	if response["text"] != "Hello WebSocket" {
		t.Errorf("Expected echo response with 'Hello WebSocket', got: %v", response)
	}
}
