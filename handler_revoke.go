package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/tobib-dev/chirpy/internal/auth"
	"github.com/tobib-dev/chirpy/internal/database"
)

func (cfg *apiConfig) handlerRevokeToken(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get bearer token", err)
		return
	}

	rToken, err := cfg.db.GetToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get refreshed tokens", err)
		return
	}

	err = cfg.db.UpdateToken(r.Context(), database.UpdateTokenParams{
		Token:     rToken.Token,
		UpdatedAt: time.Now(),
		RevokedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unable to revoke token", err)
		return
	}

	type response struct{}
	respondWithJSON(w, http.StatusNoContent, response{})
}
