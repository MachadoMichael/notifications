package whatsapp

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

var Client *twilio.RestClient
var Params *api.CreateMessageParams

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	sender := os.Getenv("SENDER")
	Client = twilio.NewRestClient()
	Params.SetFrom(sender)
}
