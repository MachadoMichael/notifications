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
		v1.POST("/send", handler.Send)
	}

	v2 := router.Group("/notifications/admin")
	{
		v2.POST("/create", handler.Create)
		v2.GET("/read", handler.Read)
		v2.PUT("/update", handler.Update)
	}

	router.Run(":" + port)
}
