package main

import (
	"net/http"

	"github.com/Sakaar-Sen/rssagg/internal/auth"
	"github.com/Sakaar-Sen/rssagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 401, err.Error())
			return
		}

		user, err := cfg.DB.GetUserByApiKey(r.Context(), apikey)
		if err != nil {
			respondWithError(w, 400, "User Not found")
			return
		}

		handler(w, r, user)
	}
}
