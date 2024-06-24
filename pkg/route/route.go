package route

import (
	"github.com/MachadoMichael/notifications/infra"
	"github.com/MachadoMichael/notifications/pkg/handler"
	"github.com/gin-gonic/gin"
)

func Init() {
	port := infra.Config.Port

	router := gin.Default()
	v1 := router.Group("/notifications/whatsapp")
	{
		v1.POST("/send", handler.Sender)
	}

	router.Run(":" + port)
}
