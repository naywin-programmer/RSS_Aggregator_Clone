package main

import (
	"log"
	"net/http"
	"os"

	cf "github.com/naywin-programmer/RSS_Aggregator/configs"
	rt "github.com/naywin-programmer/RSS_Aggregator/routes"
)

func main() {
	cf.SetupEnv()
	cf.ConnectDatabase()
	cf.RunServices()
	cf.SetUpView()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatalln("PORT didn't found in the environment setting.")
	}

	server := &http.Server{
		Handler: cf.GetRouter(rt.WebRoutes, rt.ApiRoutes),
		Addr:    ":" + portString,
	}

	log.Println("Server is running on", portString)
	serverErr := server.ListenAndServe()
	// serverErrTLS := server.ListenAndServeTLS("certFilePath", "keyFilePath")
	if serverErr != nil {
		log.Fatalln(serverErr)
	}
}
