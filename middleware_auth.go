package main

import (
	"github.com/eclipsemode/RSS-aggregator/internal/auth"
	"github.com/eclipsemode/RSS-aggregator/internal/database"
	"net/http"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		handler(w, r, user)
	}
}
