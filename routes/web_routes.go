package routes

import (
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	ct "github.com/naywin-programmer/RSS_Aggregator/controllers"
	"github.com/naywin-programmer/RSS_Aggregator/devtools"
	ut "github.com/naywin-programmer/RSS_Aggregator/utils"
)

func useDevTools(r *chi.Mux) {
	if appEnv := strings.ToLower(os.Getenv("APP_ENV")); appEnv == "dev" {
		r.HandleFunc("/ws-devtools", devtools.CreateWebSocketDevToolsServer)
	}
}

func WebRoutes(r *chi.Mux) {
	useDevTools(r)

	// === ROUTES ===
	r.NotFound(func(w http.ResponseWriter, req *http.Request) {
		ut.View404(w)
	})

	r.Get("/", ct.WelcomeController.Welcome)
	r.Get("/about-us", ct.WelcomeController.AboutUs)

}
