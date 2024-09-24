package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/naywin-programmer/RSS_Aggregator/internal/database"
)

type DatabaseQueryApiConfig struct {
	DB *database.Queries
}

func connectDatabase() DatabaseQueryApiConfig {

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatalln("DB_URL didn't found in the environment setting.")
	}

	connection, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalln("Can't connect to database.")
	}

	pingErr := connection.Ping()
	if pingErr != nil {
		log.Fatalln(pingErr)
	}
	log.Println("Database Connected")

	return DatabaseQueryApiConfig{
		DB: database.New(connection),
	}
}
