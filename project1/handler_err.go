package main

import "net/http"

// respondWithError sends an error response with the given status code and message.
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}
