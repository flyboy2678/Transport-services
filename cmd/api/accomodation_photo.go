package main

import (
	"net/http"
	"strconv"
	"transportService/internal/store"

	"github.com/go-chi/chi/v5"
)

type CreateAccomodationPhotoPayload struct {
	Accomodation_id int64  `json:"accomodation_id" validate:"required"`
	Photo_url       string `json:"photo_url" validate:"required"`
}

func (app *application) createAccomodationPhotoHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateAccomodationPhotoPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	accomodation := &store.AccomodationPhoto{
		Accomodation_id: payload.Accomodation_id,
		Photo_url:       payload.Photo_url,
	}

	ctx := r.Context()

	if err := app.store.AccomodationPhotos.Create(ctx, accomodation); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, accomodation); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getAccomodationPhotoById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	photo, err := app.store.AccomodationPhotos.GetById(ctx, id)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, photo); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getAccomodationPhotoByAccomodationId(w http.ResponseWriter, r *http.Request) {
	accomodationId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	photos, err := app.store.AccomodationPhotos.GetByAccomodationId(ctx, accomodationId)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, photos); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) deleteAccomodationPhotoById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.AccomodationPhotos.DeleteByID(ctx, id); err != nil {
		app.internalServerError(w, r, err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) deleteAccomodationPhotoByAccomodationId(w http.ResponseWriter, r *http.Request) {
	accomodationId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.AccomodationPhotos.DeleteByAccomodationId(ctx, accomodationId); err != nil {
		app.internalServerError(w, r, err)
	}

	w.WriteHeader(http.StatusNoContent)
}
