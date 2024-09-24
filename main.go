package main

import (
	"log"
	"net/http"
	"os"
)

// 7:26:15
func main() {
	setupEnv()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatalln("PORT didn't found in the environment setting.")
	}

	server := &http.Server{
		Handler: getRouter(connectDatabase()),
		Addr:    ":" + portString,
	}

	log.Println("Server is running on", portString)
	serverErr := server.ListenAndServe()
	if serverErr != nil {
		log.Fatalln(serverErr)
	}
}
