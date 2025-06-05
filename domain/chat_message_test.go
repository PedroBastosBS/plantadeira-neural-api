package entities_test

import (
	"plantadeira-neural-api/domain/Entities"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewChatMessage_Creation(t *testing.T) {
	chatMsg := entities.NewChatMessage()
	require.NotNil(t, chatMsg, "NewChatMessage should return a non-nil pointer")
}

func TestChatMessage_FieldAssignment(t *testing.T) {
	chatMsg := entities.NewChatMessage()
	id := uuid.New().String()
	userID := uuid.New().String()
	conversationID := uuid.New().String()
	content := "Hello world"
	createdAt := time.Now()

	chatMsg.ID = id
	chatMsg.UserID = userID
	chatMsg.ConversationID = conversationID
	chatMsg.Content = content
	chatMsg.CreatedAt = createdAt

	require.Equal(t, id, chatMsg.ID)
	require.Equal(t, userID, chatMsg.UserID)
	require.Equal(t, conversationID, chatMsg.ConversationID)
	require.Equal(t, content, chatMsg.Content)
	require.Equal(t, createdAt, chatMsg.CreatedAt)
}
