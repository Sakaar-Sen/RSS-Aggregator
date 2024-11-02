package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Sakaar-Sen/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apicfg *apiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Invalid request payload")
		return
	}

	user, err := apicfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		fmt.Println(err)
		respondWithError(w, 500, "Internal server error")
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apicfg *apiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apicfg *apiConfig) HandlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apicfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		respondWithError(w, 500, "Internal server error")
		return
	}

	respondWithJSON(w, 200, posts)
}
