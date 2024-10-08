package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/avar348/golangrssaggregator/internal/database"
	"github.com/avar348/golangrssaggregator/models"
	"github.com/google/uuid"
)

func (apiConfic *apiConfic) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}
	user, err := apiConfic.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt create user: %v", err))
	}

	respondWithJson(w, 200, models.DatabaseUserToUser(user))
}

func (apiConfic *apiConfic) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJson(w, 200, models.DatabaseUserToUser(user))
}

func (apiConfic *apiConfic) handleGetPostForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	post, err := apiConfic.DB.GetPostForUser(r.Context(), database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt get Post: %v", err))
	}

	respondWithJson(w, 200, models.DatabasePoststoPosts(post))

}
