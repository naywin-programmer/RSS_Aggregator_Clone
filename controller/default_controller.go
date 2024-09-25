package controller

import (
	"fmt"
	"html"
	"net/http"

	"github.com/naywin-programmer/RSS_Aggregator/utils"
)

func HelloWorldRespond(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, %q", html.EscapeString(r.URL.Path))
}

func ErrorResponse(w http.ResponseWriter, r *http.Request) {
	utils.RespondErrorWithJSON(w, 400, "Something went wrong.")
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	// all the field names in the struct must start with Uppercase. Otherwise, it will not export in the respond data
	type techstack struct {
		Title       string `json:"title"`
		Version     string `json:"version"`
		Description string `json:"description"`
		Link        string `json:"link"`
	}

	type dependency struct {
		Name        string `json:"title"`
		Version     string `json:"version"`
		Description string `json:"description"`
		Link        string `json:"link"`
	}

	type data struct {
		Status       string       `json:"status"`
		Developer    string       `json:"developer"`
		AppName      string       `json:"app_name"`
		Language     string       `json:"language"`
		Techstacks   []techstack  `json:"techstacks"`
		Dependencies []dependency `json:"dependencies"`
	}

	utils.RespondWithJSON(w, 200, data{
		Status:    "Active",
		Developer: "Nay Win",
		AppName:   "RSS Aggregator",
		Language:  "Golang",
		Techstacks: []techstack{
			{Title: "Golang", Version: "1.23", Description: "Server-side Programming Language", Link: ""},
			{Title: "HTTP Server", Version: "v5.1.0", Description: "Chi", Link: "github.com/go-chi/chi/v5"},
		},
		Dependencies: []dependency{
			{Name: "Chi CORS", Version: "v1.2.1", Description: "CORS", Link: "github.com/go-chi/cors"},
			{Name: "GoDotenv", Version: "1.12", Description: "For .env file", Link: "github.com/joho/godotenv"},
		},
	})
}
