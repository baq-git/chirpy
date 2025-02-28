package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type parameters struct {
	Body string `json:"body"`
}
type returnVals struct {
	Valid       bool   `json:"valid"`
	CleanedBody string `json:"cleaned_body"`
}

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		responseWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	words := strings.Split(params.Body, " ")
	for i, w := range words {
		if strings.ToLower(w) == "kerfuffle" || strings.ToLower(w) == "sharbert" || strings.ToLower(w) == "fornax" {
			words[i] = "****"
		}
	}
	cleanStrings := strings.Join(words, " ")

	responseWithJson(w, http.StatusOK, returnVals{
		Valid:       true,
		CleanedBody: cleanStrings,
	})
}
