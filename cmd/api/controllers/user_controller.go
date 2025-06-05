package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"plantadeira-neural-api/application/services"
	"plantadeira-neural-api/domain"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

// CreateUser godoc
// @Summary      Cria um novo usuário
// @Description  Cria um usuário com nome, email e id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      domain.User  true  "Usuário a ser criado"
// @Success      200   {object}  domain.User
// @Failure      400   {object}  map[string]string
// @Router       /users [post]
func (c *UserController) CreateUser(ctx *gin.Context) {
	var user domain.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Payload invalido"})
		return
	}

	if err := c.userService.Save(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Usuario criado com sucesso", "user": user})
}
