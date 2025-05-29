package main

import (
	"encoding/json"
	"net/http"
)

func handlerValidate(w http.ResponseWriter, r *http.Request) {

	type chirpResponse struct {
		Body string `json:"body"`
	}

	type returnError struct {
		Error string `json:"error"`
	}

	type returnValid struct {
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(r.Body)
	chirp := chirpResponse{}
	err := decoder.Decode(&chirp)
	if err != nil {
		respBody := returnError{
			Error: "Something went wrong",
		}

		data, err := json.Marshal(respBody)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(data)
	}

	if len(chirp.Body) > 140 {
		respBody := returnError{
			Error: "Chirp is too long",
		}

		data, err := json.Marshal(respBody)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(data)
		return
	}

	respBody := returnValid{
		Valid: true,
	}

	data, err := json.Marshal(respBody)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)

}
