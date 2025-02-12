package main

import (
	"net/http"
	"strconv"
	"transportService/internal/store"

	"github.com/go-chi/chi/v5"
)

type CreateSubPayload struct {
	UserID int64  `json:"user_id" validate:"required"`
	Email  string `json:"email" validate:"required"`
}

// CreateSubscription godoc
//
// @Summary Creates a subscription
// @Description Creates a subscription
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param payload body	 CreateSubPayload		true	"Post payload"
//
//	@Success		202		{object}	store.Subscription
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/subscriptions [post]
func (app *application) createSubHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateSubPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	sub := &store.Subscription{
		User_id: payload.UserID,
		Email:   payload.Email,
	}

	ctx := r.Context()

	if err := app.store.Subscriptions.Create(ctx, sub); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, sub); err != nil {
		app.internalServerError(w, r, err)
		return
	}

}

// GetAllSubscriptions godoc
//
// @Summary Fetches all subscriptions
// @Description Fetches all subscriptions
// @Tags subscriptions
// @Accept json
// @Produce json
//
//	@Success		200	{object}	store.Subscription
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/subscriptions [get]
func (app *application) getAllSubsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	subs, err := app.store.Subscriptions.GetAll(ctx)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, subs); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// DeleteUserByEmail godoc
//
// @Summary Deletes a subscription
// @Description Deletes a subscription by email
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param email	path		string	true	"Subscription email"
//
//	@Success		204	{object} string
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/subscriptions/email/{email} [delete]
func (app *application) deleteSubByEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	ctx := r.Context()

	if err := app.store.Subscriptions.DeleteByEmail(ctx, email); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteUserById godoc
//
// @Summary Deletes a subscription
// @Description Deletes a subscription by id
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id	path		int	true	"Subscription ID"
//
//	@Success		204	{object} string
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/subscriptions/id/{id} [delete]
func (app *application) deleteSubByUserIdHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Subscriptions.DeleteByUserID(ctx, userID); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
