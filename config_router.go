package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func getRouter(dbQueryApiConfig DatabaseQueryApiConfig) *chi.Mux {

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

	webRoutes(router, dbQueryApiConfig)

	apiRouter := chi.NewRouter()
	apiRoutes(apiRouter, dbQueryApiConfig)

	router.Mount("/v1", apiRouter)

	return router

}
