package repositories

import (
	"gorm.io/gorm"
	"plantadeira-neural-api/domain"
	"time"
)

type ChatMessageRepository struct {
	db *gorm.DB
}

func NewChatMessageRepository(db *gorm.DB) *ChatMessageRepository {
	return &ChatMessageRepository{db: db}
}

func (r *ChatMessageRepository) Save(message *domain.ChatMessage) error {

	if message.CreatedAt.IsZero() {
		message.CreatedAt = time.Now()
	}

	return r.db.Create(message).Error
}
