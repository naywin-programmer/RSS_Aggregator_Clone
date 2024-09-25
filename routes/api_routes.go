package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/naywin-programmer/RSS_Aggregator/controller"
)

func ApiRoutes(router *chi.Mux) {

	router.Post("/create-user", controller.ResponseCreateUser)
	router.Get("/auth", controller.ResponseGetUserByApiKey)

	router.Get("/healthz", controller.Healthz)
	router.Get("/error", controller.ErrorResponse)

}
