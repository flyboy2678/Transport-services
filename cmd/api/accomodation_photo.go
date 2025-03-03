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

// CreatePhoto godoc
//
// @Summary Creates a photo
// @Description Creates a photo
// @Tags accomodationPhotos
// @Accept json
// @Produce json
// @Param payload body	 CreateAccomodationPhotoPayload		true	"Post payload"
//
//	@Success		202		{object}	store.AccomodationPhoto
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/accomodationPhotos [post]
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

// GetPhotoById godoc
//
// @Summary Fetches a photo by id
// @Description Fetches a photo by id
// @Tags accomodationPhotos
// @Accept json
// @Produce json
// @Param id path int true "photo id"
//
//	@Success		200	{object}	store.AccomodationPhoto
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/accomodationPhotos/id/{id} [get]
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

// GetPhotoByTripId godoc
//
// @Summary Fetches photo by a trip id
// @Description Fetches photo by a trip id
// @Tags accomodationPhotos
// @Accept json
// @Produce json
// @Param id path int true "Accomodation id"
//
//	@Success		200	{object}	store.AccomodationPhoto
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/accomodationPhotos/accomodation/{id} [get]
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

// DeletePhotosById godoc
//
// @Summary Deletes a photo
// @Description Deletes a photo by id
// @Tags accomodationPhotos
// @Accept json
// @Produce json
// @Param id	path		int	true	"photo ID"
//
//	@Success		204	{object} string
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/accomodationPhotos/id/{id} [delete]
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

// DeletePhotosByAccomodationId godoc
//
// @Summary Deletes photos
// @Description Deletes photos by a accomodation id
// @Tags accomodationPhotos
// @Accept json
// @Produce json
// @Param id	path		int	true	"accomodation ID"
//
//	@Success		204	{object} string
//	@Failure		404	{object}	error
//	@Router			/accomodationPhotos/accomodationId/{id} [delete]
//	@Failure		500	{object}	error
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
