package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, status int, msg string) {
	if status > 499 {
		log.Printf("Server error: %s", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, status, errResponse{Error: msg})
}

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
