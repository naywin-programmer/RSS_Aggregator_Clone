package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func toJSON(payload interface{}) ([]byte, bool) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Fail to convert to JSON format: %v", payload)
		return nil, true
	}

	return data, false
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := toJSON(payload)
	if err {
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondErrorWithJSON(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Server 5XX Error:", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}
