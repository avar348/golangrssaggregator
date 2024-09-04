package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/avar348/golangrssaggregator/internal/database"
	"github.com/avar348/golangrssaggregator/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiConfic *apiConfic) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}
	feedFollows, err := apiConfic.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt create feed: %v", err))
	}

	respondWithJson(w, 200, models.DatabaseFeedFollowtoFeedFollow(feedFollows))
}

func (apiConfic *apiConfic) handleGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiConfic.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt create feed: %v", err))
	}

	respondWithJson(w, 200, models.DatabaseFeedFollowstoFeedFollows(feedFollows))
}

func (apiConfic *apiConfic) handleDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdSTR := chi.URLParam(r, "id")

	feedFollowId, err := uuid.Parse(feedFollowIdSTR)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt parse uuid: %v", err))
	}

	err = apiConfic.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt delete feed follow: %v", err))
	}

	respondWithJson(w, 200, struct{}{})
}
