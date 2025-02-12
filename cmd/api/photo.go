package main

import (
	"net/http"
	"strconv"
	"transportService/internal/store"

	"github.com/go-chi/chi/v5"
)

type CreatePhotoPayload struct {
	Trip_id   int64  `json:"trip_id" validate:"required"`
	Photo_url string `json:"photo_url" validate:"required"`
}

// CreatePhoto godoc
//
// @Summary Creates a photo
// @Description Creates a photo
// @Tags photos
// @Accept json
// @Produce json
// @Param payload body	 CreatePhotoPayload		true	"Post payload"
//
//	@Success		202		{object}	store.Photo
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/photo [post]
func (app *application) createPhotoHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePhotoPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	photo := &store.Photo{
		Trip_id:   payload.Trip_id,
		Photo_url: payload.Photo_url,
	}

	ctx := r.Context()

	if err := app.store.Photos.Create(ctx, photo); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, photo); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetPhotoBygodoc
//
// @Summary Fetches a photo by id
// @Description Fetches a photo by id
// @Tags photos
// @Accept json
// @Produce json
// @Param id path int true "Photo id"
//
//	@Success		200	{object}	store.Photo
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/photos/id/{id} [get]
func (app *application) getPhotoByIdHandler(w http.ResponseWriter, r *http.Request) {
	photoId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	photo, err := app.store.Photos.GetByID(ctx, photoId)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, photo); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetPhotosByTripId godoc
//
// @Summary Fetches photos by a trip id
// @Description Fetches photos by a trip id
// @Tags photos
// @Accept json
// @Produce json
// @Param id path int true "Trip id"
//
//	@Success		200	{object}	store.Photo
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/photos/tripId/{id} [get]
func (app *application) getPhotosByTripIdHandler(w http.ResponseWriter, r *http.Request) {
	tripId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	photos, err := app.store.Photos.GetByID(ctx, tripId)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, photos); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// DeletePhotosByTripId godoc
//
// @Summary Deletes a photo
// @Description Deletes a photo by id
// @Tags photos
// @Accept json
// @Produce json
// @Param id	path		int	true	"photo ID"
//
//	@Success		204	{object} string
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/photos/id/{id} [delete]
func (app *application) DeletePhotoByIdHandler(w http.ResponseWriter, r *http.Request) {
	photoId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Photos.DeleteByID(ctx, photoId); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletePhotosByTripId godoc
//
// @Summary Deletes photos
// @Description Deletes photos by a trip id
// @Tags photos
// @Accept json
// @Produce json
// @Param id	path		int	true	"Trip ID"
//
//	@Success		204	{object} string
//	@Failure		404	{object}	error
//	@Router			/photos/tripId/{id} [delete]
//	@Failure		500	{object}	error
func (app *application) DeletePhotosByTripHandler(w http.ResponseWriter, r *http.Request) {
	tripId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Photos.DeleteByTripID(ctx, tripId); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
