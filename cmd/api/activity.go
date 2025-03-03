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
