package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/tobib-dev/chirpy/internal/auth"
	"github.com/tobib-dev/chirpy/internal/database"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type paramaters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	type response struct {
		User
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	params := paramaters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters", err)
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	err = auth.CheckPasswordHash(user.HashedPassword, params.Password)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	expirationAT := time.Duration(time.Hour)
	tokenString, err := auth.MakeJWT(user.ID, cfg.serverToken, expirationAT)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect secret token", err)
		return
	}

	rTokenString, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "We hit a snag", err)
		return
	}

	expirationRT := time.Now().Add(time.Hour * 24 * 60)
	err = cfg.db.SaveToken(r.Context(), database.SaveTokenParams{
		Token:     rTokenString,
		UserID:    user.ID,
		ExpiresAt: expirationRT,
		RevokedAt: sql.NullTime{},
	})

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		},
		Token:        tokenString,
		RefreshToken: rTokenString,
	})
}
