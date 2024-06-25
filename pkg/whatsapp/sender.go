package whatsapp

import (
	"encoding/json"

	"github.com/MachadoMichael/notifications/infra/database"
	"github.com/MachadoMichael/notifications/pkg/logger"
	"github.com/MachadoMichael/notifications/schema"
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
	Params.SetTo(recipient)
	Params.SetBody(body)

	res, err := Client.Api.CreateMessage(Params)
	if err != nil {
		return err
	} else {
		if res.Body != nil {
			println(*res.Body)
			return nil
		} else {
			println(res.Body)
			return nil
		}
	}

}
