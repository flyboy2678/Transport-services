package main

import (
	"net/http"
	"transportService/internal/store"

	"github.com/go-chi/chi/v5"
)

type CreateInvoicePayload struct {
	Payment_id     int64  `json:"payment_id" validate:"required"`
	Invoice_number string `json:"invoice_number" validate:"required"`
	Due_date       string `json:"due_date" validate:"required"`
	Status         string `json:"status" validate:"required"`
}

// CreateInvoice godoc
//
// @Summary Creates a invoice
// @Description Creates a invoice
// @Tags invoices
// @Accept json
// @Produce json
// @Param payload body	 CreateInvoicePayload		true	"Post payload"
//
//	@Success		202		{object}	store.Invoice
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/invoice [post]
func (app *application) createInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateInvoicePayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	invoice := &store.Invoice{
		Payment_id:     payload.Payment_id,
		Invoice_number: payload.Invoice_number,
		Due_date:       payload.Due_date,
		Status:         payload.Status,
	}

	ctx := r.Context()

	if err := app.store.Invoices.Create(ctx, invoice); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, invoice); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// getInoiceByInvoiceNumber godoc
//
// @summary fetches a invoice
// @description fetches a invoice by invoice number
// @tags invoices
// @accept json
// @produce json
// @param invoiceNumber path string true "invoice number"
//
//	@success		200	{object}	store.Invoice
//	@failure		404	{object}	error
//	@failure		500	{object}	error
//	@router			/invoice/invoiceNumber/{invoiceNumber} [get]
func (app *application) getInvoiceByInvoiceNumberHandler(w http.ResponseWriter, r *http.Request) {
	invoiceNumber := chi.URLParam(r, "invoiceNumber")

	ctx := r.Context()

	invoice, err := app.store.Invoices.GetByInvoiceNumber(ctx, invoiceNumber)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, invoice); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type UpdateInvoicePayload struct {
	InvoiceNumber string `json:"invoice_number" validate:"required"`
	Status        string `json:"status" validate:"required"`
}

// UpdateInvoice godoc
//
// @Summary		Updates a invoice
// @Description	Updates a invoice by invoice number
// @Tags			invoices
// @Accept			json
// @Produce		json
// @Param			invoiceNumber		path		string					true	"invoice number"
// @Param			payload	body		UpdateInvoicePayload	true	"Post payload"
//
//	@Success		200		{object}	store.Invoice
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/invoice/invoiceNumber/{invoiceNumber} [patch]
func (app *application) updateInvoiceByInvoiceNumberHandler(w http.ResponseWriter, r *http.Request) {
	var payload UpdateInvoicePayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	invoice := &store.Invoice{
		Invoice_number: payload.InvoiceNumber,
		Status:         payload.Status,
	}

	ctx := r.Context()

	if err := app.store.Invoices.UpdateByInvoiceNumber(ctx, invoice); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, invoice); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
