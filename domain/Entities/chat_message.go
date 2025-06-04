package Entities

import "time"

type ChatMessage struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	ConversationID string    `json:"conversation_id"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
}

func NewChatMessage() *ChatMessage {
	return &ChatMessage{}
}
