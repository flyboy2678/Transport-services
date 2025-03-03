package main

import (
	"net/http"
	"strconv"
	"transportService/internal/store"

	"github.com/go-chi/chi/v5"
)

type CreateActivityPayload struct {
	Trip_id     int64   `json:"trip_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

// CreateActvity godoc
//
// @Summary Creates a activity
// @Description Creates a activity
// @Tags activities
// @Accept json
// @Produce json
// @Param payload body	 CreateActivityPayload		true	"Post payload"
//
//	@Success		202		{object}	store.Activity
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/activity [post]
func (app *application) createActivityHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateActivityPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	activity := &store.Activity{
		Trip_id:     payload.Trip_id,
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
	}

	ctx := r.Context()

	if err := app.store.Activities.Create(ctx, activity); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, activity); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetActivityBygodoc
//
// @Summary Fetches a activity by id
// @Description Fetches a activity by id
// @Tags activities
// @Accept json
// @Produce json
// @Param id path int true "Activity id"
//
//	@Success		200	{object}	store.Activity
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/activities/id/{id} [get]
func (app *application) getActivityByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	activity, err := app.store.Activities.GetById(ctx, id)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, activity); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetActivitiesByTripId godoc
//
// @Summary Fetches activities by a trip id
// @Description Fetches activities by a trip id
// @Tags activities
// @Accept json
// @Produce json
// @Param id path int true "Trip id"
//
//	@Success		200	{object}	store.Activity
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/activities/tripId/{id} [get]
func (app *application) getActivityByTripIdHandler(w http.ResponseWriter, r *http.Request) {
	tripId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	activitys, err := app.store.Activities.GetById(ctx, tripId)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, activitys); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type UpdateActivityPayload struct {
	ID          int64   `json:"id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

func (app *application) updateActivityByID(w http.ResponseWriter, r *http.Request) {
	var payload UpdateActivityPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	activity := &store.Activity{
		ID:          payload.ID,
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
	}

	ctx := r.Context()

	if err := app.store.Activities.UpdateById(ctx, activity); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, activity); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
