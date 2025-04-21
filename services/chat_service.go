package services

import (
	"github.com/harsh6373/chat-app-go/config"
	"github.com/harsh6373/chat-app-go/models"
)

func SaveMessage(senderID, receiverID uint, content string) (*models.Message, error) {
	message := &models.Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		Read:       false,
	}

	if err := config.DB.Create(message).Error; err != nil {
		return nil, err
	}

	return message, nil
}

func GetChatHistory(user1, user2 uint) ([]models.Message, error) {
	var messages []models.Message

	if err := config.DB.
		Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", user1, user2, user2, user1).
		Order("created_at asc").
		Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}
