package main

import (
	"net/http"
	"strconv"
	"transportService/internal/store"

	"github.com/go-chi/chi/v5"
)

type CreateAccomodationPayload struct {
	Trip_id         int64   `json:"trip_id" validate:"required"`
	Name            string  `json:"name" validate:"required"`
	Description     string  `json:"description" validate:"required"`
	Price_per_night float64 `json:"price_per_night" validate:"required"`
}

func (app *application) createAccomodationHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateAccomodationPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	accomodation := &store.Accomodation{
		Trip_id:         payload.Trip_id,
		Name:            payload.Name,
		Decription:      payload.Description,
		Price_per_night: payload.Price_per_night,
	}

	ctx := r.Context()

	if err := app.store.Accomodations.Create(ctx, accomodation); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, accomodation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getAccomodationByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	accomodation, err := app.store.Accomodations.GetByID(ctx, id)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, accomodation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getAccomodationByTripIdHandler(w http.ResponseWriter, r *http.Request) {
	tripId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	accomodations, err := app.store.Accomodations.GetByID(ctx, tripId)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, accomodations); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type UpdateAccomodationPayload struct {
	ID              int64   `json:"id" validate:"required"`
	Name            string  `json:"name" validate:"required"`
	Description     string  `json:"description" validate:"required"`
	Price_per_night float64 `json:"price_per_night" validate:"required"`
}

func (app *application) updateAccomodationByID(w http.ResponseWriter, r *http.Request) {
	var payload UpdateAccomodationPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	accomodation := &store.Accomodation{
		ID:              payload.ID,
		Name:            payload.Name,
		Decription:      payload.Description,
		Price_per_night: payload.Price_per_night,
	}

	ctx := r.Context()

	if err := app.store.Accomodations.UpdateByID(ctx, accomodation); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, accomodation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
