package entities

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type ChatMessage struct {
	ID             string    `json:"id" valid:"uuid"`
	UserID         string    `json:"user_id" valid:"uuid"`
	ConversationID string    `json:"conversation_id" valid:"-"`
	Content        string    `json:"content" valid:"-"`
	CreatedAt      time.Time `json:"created_at" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewChatMessage() *ChatMessage {
	return &ChatMessage{}
}

func (c *ChatMessage) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}
	return nil
}
