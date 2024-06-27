package whatsapp

import (
	"encoding/json"
	"errors"

	"github.com/MachadoMichael/notifications/infra"
	"github.com/MachadoMichael/notifications/infra/database"
	"github.com/MachadoMichael/notifications/pkg/logger"
	"github.com/MachadoMichael/notifications/schema"
	"github.com/go-resty/resty/v2"
	"golang.org/x/exp/slog"
)

func Sender(enterpriseName, recipient, body string) error {

	enterpriseData, err := database.EnterpriseRepo.ReadOne(enterpriseName)
	if err != nil {
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
	}

	var enterprise schema.Enterprise
	json.Unmarshal([]byte(enterpriseData), &enterprise)

	apiUrl := "https://graph.facebook.com/v13.0/" + enterprise.PhoneNumberId + "/messages"
	client := resty.New()

	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                "recipient",
		"type":              "text",
		"text": map[string]string{
			"body": body,
		},
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+infra.Config.ApiToken).
		SetBody(payload).
		Post(apiUrl)

	if err != nil {
		logger.ErrorLogger.Write(slog.LevelError, err.Error())
		return err
	}

	if resp.IsError() {
		logger.ErrorLogger.Write(slog.LevelError, "Whatsapp response error")
		return errors.New("Whatsapp response error")
	}

	logger.MessageLogger.Write(slog.LevelInfo, "Message sent successfully.")

	return nil
}
