package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func SetupEnv() {
	filename := `.env`

	envErr := godotenv.Load(filename)
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
}
