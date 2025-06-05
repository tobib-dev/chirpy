package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUpgradeMembership(w http.ResponseWriter, r *http.Request) {
	type Data struct {
		UserID string `json:"user_id"`
	}

	type parameters struct {
		Event string `json:"event"`
		Data  Data   `json:"data"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "couldn't decode parameters", err)
		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	uID, err := uuid.Parse(params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get user id", err)
		return
	}

	_, err = cfg.db.UpgradeToChirpyRed(r.Context(), uID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "couldn't find user", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
