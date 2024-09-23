package main

import (
	"fmt"
	"html"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func webRoutes(router *chi.Mux) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World, %q", html.EscapeString(r.URL.Path))
	})
}

func apiRoutes(router *chi.Mux) {
	router.Get("/healthz", healthz)
	router.Get("/error", errorResponse)
}
