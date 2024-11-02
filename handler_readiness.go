package main

import "net/http"

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}

func HandlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "Ran into an error")
}
