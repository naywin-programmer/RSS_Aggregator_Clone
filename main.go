package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load(`.env`)
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatalln("PORT didn't found in the environment setting.")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	webRoutes(router)

	apiRouter := chi.NewRouter()
	apiRoutes(apiRouter)

	router.Mount("/v1", apiRouter)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Println("Server is running on", portString)
	serverErr := server.ListenAndServe()
	if serverErr != nil {
		log.Fatalln(serverErr)
	}
}
