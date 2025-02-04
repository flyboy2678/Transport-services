package main

import (
	"net/http"
)

type CreateUserPayload struct {
	First_name string `json:"first_name" validate:"required,max=100"`
	Last_name  string `json:"last_name" validate:"required,max=100"`
	Email      string `json:"email" validate:"required,max=255"`
	Password   string `json:"password" validate:"required,min=8,max=24"`
	Phone      string `json:"phone" validate:"required,max=255"`
}

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserPayload

	if err := readJSON(w, r, &payload); err != nil {
		// app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		// app.badRequestResponse(w, r, err)
		return
	}

}
