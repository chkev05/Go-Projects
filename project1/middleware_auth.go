package main

import (
	"fmt"
	"net/http"

	"github.com/chkev05/Go-Projects/project1/internal/auth"
	"github.com/chkev05/Go-Projects/project1/internal/database"
)

type authhandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(next authhandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Unauthorized error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("User not found: %v", err))
			return
		}

		next(w, r, user)
	}
}
