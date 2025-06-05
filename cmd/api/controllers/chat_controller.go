package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"plantadeira-neural-api/application/services"
	"plantadeira-neural-api/domain"
)

type ChatController struct {
	openIaService *services.OpenIaService
}

func NewChatController(service *services.OpenIaService) *ChatController {
	return &ChatController{openIaService: service}
}

// SendMessage godoc
// @Summary      Envia uma mensagem para o Dify
// @Description  Recebe uma mensagem do usuário, envia para o Dify e retorna a resposta com possível conversation_id
// @Tags         chat
// @Accept       json
// @Produce      json
// @Param        chatMessage  body      domain.ChatMessage  true  "Mensagem do usuário"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /chat/send [post]
func (cc *ChatController) SendMessage(ctx *gin.Context) {
	var chatMessage domain.ChatMessage

	if err := ctx.ShouldBindJSON(&chatMessage); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	resp, err := cc.openIaService.SendMessage(&chatMessage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
