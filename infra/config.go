package infra

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigData struct {
	Port        string
	Account     string
	AuthToken   string
	Sender      string
	LogFilePath string
}

var Config *ConfigData

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Config = &ConfigData{
		Port:        os.Getenv("PORT"),
		Account:     os.Getenv("ACCOUNT"),
		AuthToken:   os.Getenv("AUTH_TOKEN"),
		Sender:      os.Getenv("SENDER"),
		LogFilePath: os.Getenv("LOG_FILE_PATH"),
	}

}
