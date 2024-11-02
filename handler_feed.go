package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Sakaar-Sen/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apicfg *apiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Invalid request payload")
		return
	}

	feed, err := apicfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		fmt.Println(err)
		respondWithError(w, 500, "Could not create feed")
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apicfg *apiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feed, err := apicfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 500, "Could not fetch feeds")
		return
	}

	respondWithJSON(w, 200, databaseFeedsToFeeds(feed))
}
