package handler

import (
	"net/http"

	"github.com/MachadoMichael/notifications/infra/database"
	"github.com/MachadoMichael/notifications/pkg/logger"
	"github.com/MachadoMichael/notifications/schema"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func Create(ctx *gin.Context) {
	request := schema.Enterprise{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	enterprise, err := database.EnterpriseRepo.ReadOne(request.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	if enterprise != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "This email is already in use"})
		logger.ErrorLogger.Write(slog.LevelError, "This email is already iin use")
		return
	}

	errDb := database.EnterpriseRepo.Create(request)
	if errDb != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errDb.Error})
		logger.ErrorLogger.Write(slog.LevelError, errDb.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Enterprise created successfully"})
}
