package main

import (
	"github.com/go-chi/chi/v5"
)

func webRoutes(router *chi.Mux, dbQueryApiConfig DatabaseQueryApiConfig) {
	router.Get("/", helloWorldRespond)
}

func apiRoutes(router *chi.Mux, dbQueryApiConfig DatabaseQueryApiConfig) {
	router.Post("/create-user", dbQueryApiConfig.responseCreateUser)
	router.Get("/healthz", healthz)
	router.Get("/error", errorResponse)
}
