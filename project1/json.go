package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Error response function to handle errors in a consistent way.
func respondWithError(w http.ResponseWriter, status int, msg string) {
	if status > 499 {
		log.Printf("Server error: %s", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, status, errResponse{Error: msg})
}

// respondWithJSON is a utility function to send JSON responses gets passed in a model for the data field.
func respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	dat, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Error marshalling JSON: %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dat)
}
