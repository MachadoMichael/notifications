package handler

import (
	"net/http"

	"github.com/MachadoMichael/notifications/infra/database"
	"github.com/MachadoMichael/notifications/pkg/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func Delete(ctx *gin.Context) {
	email := ctx.Param("email")

	if !isValidToken(ctx) {
		return
	}

	enterprise, errRead := database.EnterpriseRepo.ReadOne(email)
	if errRead != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errRead.Error})
		logger.ErrorLogger.Write(slog.LevelError, errRead.Error())
		return
	}

	if enterprise == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "There isn't register for this email."})

		logger.ErrorLogger.Write(slog.LevelError, "There isn't register for this email.")
		return
	}

	rows, err := database.EnterpriseRepo.Delete(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Enterprise deleted successfully", "rows_affcteds": rows})
}
