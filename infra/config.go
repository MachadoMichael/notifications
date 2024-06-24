package infra

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigData struct {
	Port      int
	Account   string
	AuthToken string
	Sender    string
}

var Config *ConfigData

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
		panic("Cannot read envPort")
	}

	Config = &ConfigData{
		Port:      port,
		Account:   os.Getenv("ACCOUNT"),
		AuthToken: os.Getenv("AUTH_TOKEN"),
		Sender:    os.Getenv("SENDER"),
	}

}
