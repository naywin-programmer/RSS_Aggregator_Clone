package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/naywin-programmer/RSS_Aggregator/internal/database"
)

//7:31
func (dbQueryApiConfig *DatabaseQueryApiConfig) responseCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondErrorWithJSON(w, 400, fmt.Sprintf("Request to JSON Parsing Error: %v", err))
		return
	}

	user, err := dbQueryApiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondErrorWithJSON(w, 400, fmt.Sprintf("Couldn't create new user: %v", err))
		return
	}

	respondWithJSON(w, 200, user)
}
