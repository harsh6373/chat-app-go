// models/message.go
package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	SenderID   uint
	ReceiverID uint
	Content    string
	Read       bool
}

type IncomingMessage struct {
	Text     string `json:"text"`
	Receiver string `json:"receiver"`
	Typing   bool   `json:"typing"` // for typing indicator
}
