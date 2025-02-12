package main

import (
	"net/http"
	"strconv"
	"transportService/internal/store"

	"github.com/go-chi/chi/v5"
)

type CreatePaymentPayload struct {
	Booking_id     int64   `json:"booking_id" validate:"required"`
	User_id        int64   `json:"user_id" validate:"required"`
	Amount         float64 `json:"amount" validate:"required"`
	Status         string  `json:"status" validate:"required"`
	Transaction_id string  `json:"transation_id" validate:"required"`
}

// CreatePayment godoc
//
// @Summary Creates a payment
// @Description Creates a payment
// @Tags payments
// @Accept json
// @Produce json
// @Param payload body	 CreatePaymentPayload		true	"Post payload"
//
//	@Success		202		{object}	store.Payment
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/payments [post]
func (app *application) createPaymentHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePaymentPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	payment := &store.Payment{
		Booking_id:     payload.Booking_id,
		User_id:        payload.User_id,
		Amount:         payload.Amount,
		Status:         payload.Status,
		Transaction_id: payload.Transaction_id,
	}

	ctx := r.Context()

	if err := app.store.Payments.Create(ctx, payment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, payment); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetPaymentByUserId godoc
//
// @Summary Fetches payments
// @Description Fetches payments by a user id
// @Tags payments
// @Accept json
// @Produce json
// @Param id path int true "payments id"
//
//	@Success		200	{object}	store.Payment
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/payments/userId/{id} [get]
func (app *application) getPaymentsByUserIdHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	payments, err := app.store.Payments.GetByUserID(ctx, userId)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, payments); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type UpdatePaymentPayload struct {
	ID     int64  `json:"id" validate:"required"`
	Status string `json:"status" validate:"required"`
}

// UpdatePayment godoc
//
// @Summary		Updates a payment
// @Description	Updates a payment by ID
// @Tags			payment
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"Payment ID"
// @Param			payload	body		UpdatePaymentPayload	true	"Post payload"
//
//	@Success		200		{object}	store.Payload
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/payments/id/{id} [patch]
func (app *application) updatePaymentByIdHandler(w http.ResponseWriter, r *http.Request) {
	var payload UpdatePaymentPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	payment := &store.Payment{
		ID:     payload.ID,
		Status: payload.Status,
	}

	ctx := r.Context()

	if err := app.store.Payments.UpdateByID(ctx, payment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, payment); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
