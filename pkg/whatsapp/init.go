package whatsapp

import (
	"github.com/MachadoMichael/notifications/infra"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

var (
	Client *twilio.RestClient
	Params *api.CreateMessageParams
)

func Init() {
	sender := infra.Config.Sender
	Client = twilio.NewRestClient()
	Params.SetFrom(sender)
}
