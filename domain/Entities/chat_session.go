package Entities

import "time"

type ChatSession struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewChatSession() *ChatSession {
	return &ChatSession{}
}
