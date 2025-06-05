// @title           Plantadeira Neural API
// @version         1.0
// @description     Esta é a API do chatbot do workshop que faz uma integração com um agente de IA
// @termsOfService  http://atto-intelligence/terms/

// @contact.name   Equipe Plantadeira
// @contact.url    http://atto-intelligence/suporte
// @contact.email  suporte@atto-intelligence.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /
package main

import (
	"github.com/gin-contrib/cors"
	"log"

	"plantadeira-neural-api/application/services"
	"plantadeira-neural-api/cmd/api/controllers"
	"plantadeira-neural-api/framework/postgres"
	"plantadeira-neural-api/framework/repositories"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "plantadeira-neural-api/docs" // Import para registrar docs do swagger
)

func main() {
	db := postgres.ConnectGORM()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get sql.DB from gorm.DB: ", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("failed to ping database: ", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	healthController := controllers.NewHealthController()
	openIaReponsitory := repositories.NewChatMessageRepository(db)
	openAiService := services.NewOpenIaService(openIaReponsitory)
	chatMessageController := controllers.NewChatController(openAiService)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", healthController.HealthCheck)

	r.POST("/users", userController.CreateUser)
	r.POST("/chat/send", chatMessageController.SendMessage)

	if err := r.Run(); err != nil {
		log.Fatal("failed to run server: ", err)
	}
}
