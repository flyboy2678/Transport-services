package main

import (
	"net/http"
	"strconv"
	"transportService/internal/store"

	"github.com/go-chi/chi/v5"
)

type CreateActiviyPhotoPayload struct {
	Activity_id int64  `json:"accomodation_id" validate:"required"`
	Photo_url   string `json:"photo_url" validate:"required"`
}

func (app *application) createActivityPhotoHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateActiviyPhotoPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	activity := &store.ActivityPhoto{
		Activity_id: payload.Activity_id,
		Photo_url:   payload.Photo_url,
	}

	ctx := r.Context()

	if err := app.store.ActivityPhotos.Create(ctx, activity); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, activity); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getActivityPhotoById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	photo, err := app.store.ActivityPhotos.GetById(ctx, id)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, photo); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getActivityPhotoByActivityId(w http.ResponseWriter, r *http.Request) {
	activityId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	photos, err := app.store.ActivityPhotos.GetById(ctx, activityId)
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

	if err := app.store.ActivityPhotos.DeleteById(ctx, id); err != nil {
		app.internalServerError(w, r, err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) deleteAccomodationPhotoByAccomodationId(w http.ResponseWriter, r *http.Request) {
	activityId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.ActivityPhotos.DeleteById(ctx, activityId); err != nil {
		app.internalServerError(w, r, err)
	}

	w.WriteHeader(http.StatusNoContent)
}
