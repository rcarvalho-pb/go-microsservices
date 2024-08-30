package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// err := app.readJson(w, r, &requestPayload)
	// if err != nil {
	// 	_ = app.errorJson(w, err, http.StatusBadRequest)
	// 	return
	// }
	body, err := io.ReadAll(r.Body)
	if err != nil {
		_ = app.errorJson(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &requestPayload); err != nil {
		_ = app.errorJson(w, err, http.StatusUnprocessableEntity)
		return
	}

	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		_ = app.errorJson(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		_ = app.errorJson(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s\n", user.Email),
		Data:    user,
	}

	app.writeJson(w, http.StatusAccepted, payload)
}
