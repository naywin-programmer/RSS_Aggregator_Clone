package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/naywin-programmer/RSS_Aggregator/controller"
)

func WebRoutes(router *chi.Mux) {
	router.Get("/", controller.HelloWorldRespond)
}
