package ws

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/harsh6373/chat-app-go/config"
	"github.com/harsh6373/chat-app-go/services"
)

type Client struct {
	Conn *websocket.Conn
	ID   string
}

type IncomingMessage struct {
	Text      string `json:"text"`
	Receiver  string `json:"receiver"`
	Typing    bool   `json:"typing"`
	Broadcast bool   `json:"broadcast"`
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
		var msg IncomingMessage
		if err := c.ReadJSON(&msg); err != nil {
			fmt.Println("Error reading JSON:", err)
			break
		}

		// Typing Indicator
		if msg.Typing && msg.Receiver != "" {
			publishTypingStatus(userID, msg.Receiver)
			continue
		}

		// Save message
		senderID, _ := strconv.Atoi(userID)
		receiverID, _ := strconv.Atoi(msg.Receiver)
		_, err := services.SaveMessage(uint(senderID), uint(receiverID), msg.Text)
		if err != nil {
			fmt.Println("Failed to save message:", err)
		}

		// Send to receiver only
		if msg.Broadcast {
			broadcast(userID, msg.Text)
		} else {
			sendToClient(msg.Receiver, fmt.Sprintf("From %s: %s", userID, msg.Text))
		}
	}

	mu.Lock()
	delete(clients, userID)
	mu.Unlock()
	fmt.Println("User disconnected:", userID)
}

func sendToClient(userID string, message string) {
	mu.Lock()
	defer mu.Unlock()
	if client, ok := clients[userID]; ok {
		client.Conn.WriteJSON(message)
	}
}

func publishTypingStatus(sender, receiver string) {
	channel := "typing:" + receiver
	payload := fmt.Sprintf("%s is typing...", sender)
	config.RedisClient.Publish(config.Ctx, channel, payload)
}

func SubscribeTypingUpdates() {
	pubsub := config.RedisClient.PSubscribe(context.Background(), "typing:*")
	go func() {
		for msg := range pubsub.Channel() {
			parts := strings.Split(msg.Channel, ":")
			if len(parts) == 2 {
				receiver := parts[1]
				sendToClient(receiver, msg.Payload)
			}
		}
	}()
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
