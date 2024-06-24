package handler

import (
	"net/http"
	"time"

	"github.com/MachadoMichael/notifications/dto"
	"github.com/MachadoMichael/notifications/infra/database"
	"github.com/MachadoMichael/notifications/pkg/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func Read(ctx *gin.Context) {
	if !isValidToken(ctx) {
		return
	}

	notifications, err := database.EnterpriseRepo.Read()
	if err != nil {
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	var content []string

	for key := range notifications {
		content = append(content, key)
	}

	layout := "2006-01-02 15:04:05"
	formattedTime := time.Now().Format(layout)

	response := dto.RegisteredEnterprisesDTO{
		Lenght:  len(content),
		Content: content,
		ReadAt:  formattedTime,
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successful to read notifications on database", "notifications": response})
}
