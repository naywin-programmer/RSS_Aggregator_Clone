package main

import (
	"log"
	"net/http"
	"os"

	"github.com/naywin-programmer/RSS_Aggregator/config"
	"github.com/naywin-programmer/RSS_Aggregator/routes"
)

func main() {
	config.SetupEnv()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatalln("PORT didn't found in the environment setting.")
	}

	config.ConnectDatabase()
	server := &http.Server{
		Handler: config.GetRouter(routes.WebRoutes, routes.ApiRoutes),
		Addr:    ":" + portString,
	}

	log.Println("Server is running on", portString)
	serverErr := server.ListenAndServe()
	if serverErr != nil {
		log.Fatalln(serverErr)
	}
}
