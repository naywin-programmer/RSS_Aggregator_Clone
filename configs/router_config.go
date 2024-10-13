package configs

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

func GetRouter(routes ...func(*chi.Mux)) *chi.Mux {

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

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	// router.Use(middleware.NoCache)
	router.Use(middleware.Throttle(50000))
	router.Use(middleware.Compress(5))
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(httprate.LimitByIP(200, time.Minute))

	createFileServer(router)

	for _, eachRoutes := range routes {
		go eachRoutes(router)
	}

	return router

}

func createFileServer(r *chi.Mux) {
	fs := http.FileServer(http.Dir("static_assets"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
}
