package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"plantadeira-neural-api/domain"
	"plantadeira-neural-api/domain/dtos"
	"plantadeira-neural-api/framework/repositories"
)

type OpenIaService struct {
	chatMessageRepo *repositories.ChatMessageRepository
}

func NewOpenIaService(chatMessageRepo *repositories.ChatMessageRepository) *OpenIaService {
	return &OpenIaService{chatMessageRepo: chatMessageRepo}
}
func (s *OpenIaService) SendMessage(chatMessage *domain.ChatMessage) (*dtos.ChatResponseDTO, error) {
	url := "https://api.dify.ai/v1/chat-messages"
	apiKey := "app-mT26Qb1ZB5WQjTMPsDe2oOkz"

	payload := map[string]interface{}{
		"inputs":        map[string]interface{}{},
		"query":         chatMessage.Content,
		"response_mode": "blocking",
		"user":          chatMessage.UserID,
	}

	if chatMessage.ConversationID != "" {
		payload["conversation_id"] = chatMessage.ConversationID
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("erro ao serializar JSON: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %v", err)
	}

	var responseDTO dtos.ChatResponseDTO

	if err := json.Unmarshal(respBody, &responseDTO); err != nil {
		return nil, fmt.Errorf("erro ao interpretar resposta JSON: %v", err)
	}

	if chatMessage.ID == "" {
		chatMessage.ID = uuid.New().String()
	}

	chatMessage.Content = responseDTO.Answer

	if chatMessage.ConversationID == "" {
		chatMessage.ConversationID = responseDTO.ConversationID
	}
	if err := s.chatMessageRepo.Save(chatMessage); err != nil {
		return nil, fmt.Errorf("erro ao salvar mensagem no banco: %v", err)
	}
	return &responseDTO, nil
}
