package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Sakaar-Sen/rssagg/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apicfg *apiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Invalid request payload")
		return
	}

	feed, err := apicfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		fmt.Println(err)
		respondWithError(w, 500, err.Error())
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feed))
}

func (apicfg *apiConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apicfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apicfg *apiConfig) HandlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedfollowid, err := uuid.Parse(feedFollowIDStr)

	if err != nil {
		respondWithError(w, 400, "Invalid feed follow ID")
		return
	}

	err = apicfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedfollowid,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	respondWithJSON(w, 200, map[string]string{"message": "Feed follow deleted successfully"})
}
