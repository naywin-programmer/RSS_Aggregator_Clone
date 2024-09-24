package main

import (
	"log"

	"github.com/joho/godotenv"
)

func setupEnv() {
	filename := `.env`

	envErr := godotenv.Load(filename)
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
}
