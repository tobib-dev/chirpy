package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/tobib-dev/chirpy/internal/database"
)

func (cfg *apiConfig) handlerGetAllChirps(w http.ResponseWriter, r *http.Request) {
	authorID := r.URL.Query().Get("author_id")
	if authorID == "" {
		chirpList, err := cfg.db.GetAllChirps(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't get all chirps", err)
			return
		}

		res := getResponseList(chirpList)
		respondWithJSON(w, http.StatusOK, res)
	} else {
		aID, err := uuid.Parse(authorID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "couldn't retrieve author id", err)
			return
		}

		chirpList, err := cfg.db.GetAllChirpsForUser(r.Context(), aID)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "couldn't find chirps", err)
			return
		}

		res := getResponseList(chirpList)
		respondWithJSON(w, http.StatusOK, res)
	}
}

func getResponseList(chirpList []database.Chirp) []Chirp {
	response := make([]Chirp, len(chirpList))
	for i, chirp := range chirpList {
		response[i] = Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		}
	}
	return response
}

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid chirp id", err)
		return
	}

	dbChirp, err := cfg.db.GetChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "couldn't find chirp", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		Body:      dbChirp.Body,
		UserID:    dbChirp.UserID,
	})
}
