package handler

import (
	"net/http"

	"github.com/MachadoMichael/notifications/pkg/logger"
	"github.com/MachadoMichael/notifications/pkg/whatsapp"
	"github.com/MachadoMichael/notifications/schema"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func Send(ctx *gin.Context) {
	request := schema.MessageRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
	}

	err := whatsapp.Sender(request.Recipient, request.Body)
	if err != nil {
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	logger.MessageLogger.Write(slog.LevelInfo, "Message sent with successfuly, from: "+request.Recipient+" & body: "+request.Body)
}
