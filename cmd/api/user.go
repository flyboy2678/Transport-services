package main

import (
	"net/http"
	"net/url"
	"strconv"
	"transportService/internal/store"

	"github.com/go-chi/chi/v5"
)

type CreateUserPayload struct {
	First_name string `json:"first_name" validate:"required,max=100"`
	Last_name  string `json:"last_name" validate:"required,max=100"`
	Email      string `json:"email" validate:"required,max=255"`
	Password   string `json:"password" validate:"required,min=8,max=24"`
	Phone      string `json:"phone" validate:"required,max=255"`
}

// CreateUser godoc
//
// @Summary Creates a user
// @Description Creates a user
// @Tags users
// @Accept json
// @Produce json
// @Param payload body	 CreateUserPayload		true	"Post payload"
//
//	@Success		202		{object}	store.User
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/users [post]
func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &store.User{
		First_name: payload.First_name,
		Last_name:  payload.Last_name,
		Email:      payload.Email,
		Password:   payload.Password,
		Phone:      payload.Phone,
	}

	ctx := r.Context()

	if err := app.store.Users.Create(ctx, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetUserById godoc
//
// @Summary Fetches a user
// @Description Fetches a user by id
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User id"
//
//	@Success		200	{object}	store.User
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/users/id/{id} [get]
func (app *application) getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	user, err := app.store.Users.GetByID(ctx, userID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetUserByEmail godoc
//
// @Summary Fetches a user
// @Description Fetches a user by id
// @Tags users
// @Accept json
// @Produce json
// @Param email path string true "User email"
//
//	@Success		200	{object}	store.User
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/users/email/{email} [get]
func (app *application) getUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	// to prevent the coversion of @ to %40
	email, err := url.QueryUnescape(chi.URLParam(r, "email"))
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	user, err := app.store.Users.GetByEmail(ctx, email)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// DeleteUserById godoc
//
// @Summary Deletes a user
// @Description Deletes a user by id
// @Tags users
// @Accept json
// @Produce json
// @Param id	path		int	true	"User ID"
//
//	@Success		204	{object} string
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/users/id/{id} [delete]
func (app *application) deleteUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Users.DeleteByID(ctx, userID); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
