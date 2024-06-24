package route

import (
	"github.com/MachadoMichael/notifications/pkg/handler"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	v1 := router.Group("/notifications/whatsapp")
	{
		v1.POST("/send", handler.Send)
	}

	router.Run()
}
