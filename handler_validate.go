package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Body string `json:"body"`
	}

	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	profane := []string{"kerfuffle", "sharbert", "fornax"}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	chirpWords := strings.Split(params.Body, " ")
	bodyList := make([]string, len(chirpWords))

	for i, word := range chirpWords {
		isProfane := false
		if slices.Contains(profane, strings.ToLower(word)) {
			isProfane = true
		}

		if isProfane {
			bodyList[i] = "****"
		} else {
			bodyList[i] = word
		}
	}

	cleanedBody := strings.Join(bodyList, " ")

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: cleanedBody,
	})
}
