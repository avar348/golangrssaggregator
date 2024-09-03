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

func (apiConfic *apiConfic) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}
	feed, err := apiConfic.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt create feed: %v", err))
	}

	respondWithJson(w, 200, models.DatabaseFeedtoFeed(feed))
}

func (apiConfic *apiConfic) getAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiConfic.DB.GetAllFeeds(r.Context())

	if err != nil {
		respondWithError(w, 400, "Couldnt retrieve feeds")
	}
	respondWithJson(w, 200, models.DatabaseFeedstoFeeds(feeds))
}

// func (apiConfic *apiConfic) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

// 	respondWithJson(w, 200, models.DatabaseUserToUser(user))
// }
