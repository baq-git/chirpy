package main

import (
	"encoding/json"
	"net/http"

	"github.com/baq-git/chirpy/internal/database"
)

type createUserParams struct {
	Email string `json:"email"`
}

type response struct {
	User database.User
}

func (cfg *apiConfig) handlerUsers(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	createUserParams := createUserParams{}
	err := decoder.Decode(&createUserParams)
	if err != nil {
		responseWithError(w, 406, "Something wrong with your email ", err)
		return
	}

	user, err := cfg.db.CreateUser(r.Context(), createUserParams.Email)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't create user", err)
		return
	}

	responseWithJson(w, http.StatusCreated, response{
		User: database.User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		},
	})
}
