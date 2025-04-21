package ws

import (
	"fmt"
	"sync"

	"github.com/gofiber/websocket/v2"
)

type Client struct {
	Conn *websocket.Conn
	ID   string
}

var clients = make(map[string]*Client)
var mu sync.Mutex

func HandleWebSocket(c *websocket.Conn) {
	userID := c.Params("userID")
	client := &Client{Conn: c, ID: userID}

	mu.Lock()
	clients[userID] = client
	mu.Unlock()

	fmt.Println("User connected:", userID)

	for {
		var msg string
		if err := c.ReadJSON(&msg); err != nil {
			fmt.Println("Error reading JSON:", err)
			break
		}
		broadcast(userID, msg)
	}

	mu.Lock()
	delete(clients, userID)
	mu.Unlock()
	fmt.Println("User disconnected:", userID)
}

func broadcast(senderID, msg string) {
	mu.Lock()
	defer mu.Unlock()
	for id, client := range clients {
		if id != senderID {
			client.Conn.WriteJSON(fmt.Sprintf("From %s: %s", senderID, msg))
		}
	}
}
