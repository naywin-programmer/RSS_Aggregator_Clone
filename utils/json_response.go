package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ToJSON(payload interface{}) ([]byte, bool) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Fail to convert to JSON format: %v", payload)
		return nil, true
	}

	return data, false
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := ToJSON(payload)
	if err {
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func RespondErrorWithJSON(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Server 5XX Error:", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	RespondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func RespondSuccessWithJSON(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Server 5XX Error:", msg)
	}

	type successResponse struct {
		Success string `json:"success"`
	}

	RespondWithJSON(w, code, successResponse{
		Success: msg,
	})
}

func DecodeJsonParams[T any](w http.ResponseWriter, r *http.Request, parameters *T) (T, error) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(parameters)
	if err != nil {
		RespondErrorWithJSON(w, http.StatusBadRequest, fmt.Sprintf("Request JSON Parsing Error: %v", err))
		return *parameters, err
	}
	return *parameters, nil
}
