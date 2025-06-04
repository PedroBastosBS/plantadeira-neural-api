package Entities

import "time"

type ChatMessage struct {
	ID        string    `json:"id"`
	SessionID string    `json:"session_id"`
	UserID    string    `json:"user_id"`
	Role      string    `json:"role"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func NewChatMessage() *ChatMessage {
	return &ChatMessage{}
}
