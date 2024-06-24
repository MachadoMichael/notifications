package route

import (
	"log"

	"github.com/MachadoMichael/notifications/infra"
	"github.com/MachadoMichael/notifications/pkg/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := infra.Config.Port

	router := gin.Default()
	v1 := router.Group("/notifications/whatsapp")
	{
		v1.POST("/send", handler.Send)
	}

	router.Run(":" + port)
}
