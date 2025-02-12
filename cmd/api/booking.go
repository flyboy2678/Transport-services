package main

import (
	"net/http"
	"strconv"
	"transportService/internal/store"

	"github.com/go-chi/chi/v5"
)

type CreateBookingPayload struct {
	User_id int64  `json:"user_id" validation:"required"`
	Trip_id int64  `json:"trip_id" validation:"required"`
	Status  string `json:"status" validation:"required"`
}

// CreateBooking godoc
//
// @Summary Creates a booking
// @Description Creates a booking
// @Tags bookings
// @Accept json
// @Produce json
// @Param payload body	 CreateBookingPayload		true	"Post payload"
//
//	@Success		202		{object}	store.Booking
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/bookings [post]
func (app *application) createBookingHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateBookingPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	booking := &store.Booking{
		User_id: payload.User_id,
		Trip_id: payload.Trip_id,
		Status:  payload.Status,
	}

	ctx := r.Context()

	if err := app.store.Bookings.Create(ctx, booking); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, booking); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetBookingById godoc
//
// @Summary Fetches a booking
// @Description Fetches a booking by id
// @Tags bookings
// @Accept json
// @Produce json
// @Param id path int true "booking id"
//
//	@Success		200	{object}	store.Booking
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/bookings/id/{id} [get]
func (app *application) getBookingByIdHandler(w http.ResponseWriter, r *http.Request) {
	bookingId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	booking, err := app.store.Bookings.GetByID(ctx, bookingId)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, booking); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetBookingByTripId godoc
//
// @Summary Fetches bookings
// @Description Fetches bookings by id
// @Tags bookings
// @Accept json
// @Produce json
// @Param id path int true "booking id"
//
//	@Success		200	{object}	store.Booking
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/bookings/tripId/{id} [get]
func (app *application) getBookingsByTripIdHandler(w http.ResponseWriter, r *http.Request) {
	tripId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	bookings, err := app.store.Bookings.GetByID(ctx, tripId)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, bookings); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetBookingByUserId godoc
//
// @Summary Fetches bookings
// @Description Fetches bookings by a user id
// @Tags bookings
// @Accept json
// @Produce json
// @Param id path int true "booking id"
//
//	@Success		200	{object}	store.Booking
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/booking/userId/{id} [get]
func (app *application) GetBookingByUserIdHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	bookings, err := app.store.Bookings.GetByID(ctx, userId)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, bookings); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type UpdateBookingPayload struct {
	ID     int64  `json:"id" validate:"required"`
	Status string `json:"status" validate:"required"`
}

// UpdateBooking godoc
//
// @Summary		Updates a booking
// @Description	Updates a booking by ID
// @Tags			bookings
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"Booking ID"
// @Param			payload	body		UpdateBookingPayload	true	"Post payload"
//
//	@Success		200		{object}	store.Booking
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/bookings/id/{id} [patch]
func (app *application) UpdateBookingByIdHandler(w http.ResponseWriter, r *http.Request) {
	var payload UpdateBookingPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	booking := &store.Booking{
		ID:     payload.ID,
		Status: payload.Status,
	}

	ctx := r.Context()

	if err := app.store.Bookings.UpdateByID(ctx, booking); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, booking); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
