package main

import (
	"net/http"
	"time"

	"github.com/tobib-dev/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefreshToken(w http.ResponseWriter, r *http.Request) {
	bearerToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't refresh tokens", err)
		return
	}

	rToken, err := cfg.db.GetToken(r.Context(), bearerToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get refresh tokens", err)
		return
	}
	if rToken.ExpiresAt.Before(time.Now()) {
		respondWithError(w, http.StatusUnauthorized, "token is expired, please update JWT", err)
		return
	}

	accessToken, err := auth.MakeJWT(rToken.UserID, cfg.serverToken, time.Duration(time.Hour))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't generate JWT", err)
		return
	}

	type response struct {
		Token string `json:"token"`
	}

	respondWithJSON(w, http.StatusOK, response{Token: accessToken})
}
