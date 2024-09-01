package main

import (
	"fmt"
	"net/http"

	"github.com/avar348/golangrssaggregator/internal/auth"
	"github.com/avar348/golangrssaggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfic) middlewareAuth(handler authedHandler) http.HandlerFunc {
	///
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Could not validate user: %v", err))
		}
		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Coudlnt get user: %v", err))
		}
		handler(w, r, user)
	}
}
