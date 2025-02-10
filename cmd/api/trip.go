package main

import (
	"net/http"
	"strconv"
	"transportService/internal/store"

	"github.com/go-chi/chi/v5"
)

type CreateTripPayload struct {
	Name            string  `json:"name" validate:"required,max=100"`
	Decription      string  `json:"description" validate:"required,max=255"`
	Location        string  `json:"location" validate:"required,max=100"`
	Start_date      string  `json:"start_date" validate:"required"`
	End_date        string  `json:"end_date" validate:"required"`
	Price           float64 `json:"price" validate:"required"`
	Seats           int     `json:"seats" validate:"required"`
	Available_seats int     `json:"available_seats" validate:"required"`
}

// CreateTrip godoc
//
// @Summary Creates a trip
// @Description Creates a trip
// @Tags trips
// @Accept json
// @Produce json
// @Param payload body	 CreateTripPayload	 true	 "Post payload"
//
//	@Success		202		{object}	store.Trip
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/trips [post]
func (app *application) createTripHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateTripPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	trip := &store.Trip{
		Name:            payload.Name,
		Decription:      payload.Decription,
		Location:        payload.Location,
		Start_date:      payload.Start_date,
		End_date:        payload.End_date,
		Price:           payload.Price,
		Seats:           payload.Seats,
		Available_seats: payload.Available_seats,
	}

	ctx := r.Context()

	if err := app.store.Trips.Create(ctx, trip); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, trip); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetAllTrips godoc
//
// @Summary Fetches all trips
// @Description Fetches all trips
// @Tags trips
// @Accept json
// @Produce json
//
//	@Success		200	{object}	store.Trip
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/trips [get]
func (app *application) getAllTripsHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	trips, err := app.store.Trips.GetAll(ctx)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, trips); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetTripById godoc
//
// @Summary Fetches a trip by id
// @Description Fetches a trip by id
// @Tags trips
// @Accept json
// @Produce json
// @Params id path int true "Trip id"
//
//	@Success		200	{object}	store.Trip
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/trips/id/{id} [get]
func (app *application) getTripByIdHandler(w http.ResponseWriter, r *http.Request) {
	tripId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	trip, err := app.store.Trips.GetByID(ctx, tripId)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, trip); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetTripByLocation godoc
//
// @Summary Fetches a trip by location
// @Description Fetches a trip by location
// @Tags trips
// @Accept json
// @Produce json
// @Param location path string true "Trip Location"
//
//	@Success		200	{object}	store.Trip
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/trips/location/{location} [get]
func (app *application) getTripsByLocationHandler(w http.ResponseWriter, r *http.Request) {
	location := chi.URLParam(r, "location")

	ctx := r.Context()

	trip, err := app.store.Trips.GetByLocation(ctx, location)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, trip); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetUpcomingTrips godoc
//
// @Summary Fetches all upcoming trips
// @Description Fetches all upcoming trips
// @Tags trips
// @Accept json
// @Produce json
//
//	@Success		200	{object}	store.Trip
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/trips/upcoming [get]
func (app *application) getUpcomingTripsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	trips, err := app.store.Trips.GetUpcoming(ctx)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, trips); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
