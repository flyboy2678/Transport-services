package main

import (
	"net/http"
	"strconv"
	"transportService/internal/store"

	"github.com/go-chi/chi/v5"
)

type CreateActivityPhotoPayload struct {
	Activity_id int64  `json:"accomodation_id" validate:"required"`
	Photo_url   string `json:"photo_url" validate:"required"`
}

// CreateActvityPhoto godoc
//
// @Summary Creates a activity photo
// @Description Creates a activity photo
// @Tags activityPhotos
// @Accept json
// @Produce json
// @Param payload body	 CreateActivityPhotoPayload		true	"Post payload"
//
//	@Success		202		{object}	store.ActivityPhoto
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/activityPhotos [post]
func (app *application) createActivityPhotoHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateActivityPhotoPayload

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

// GetActivityPhotoByIdgodoc
//
// @Summary Fetches a activity by id
// @Description Fetches a activity by id
// @Tags activityPhotos
// @Accept json
// @Produce json
// @Param id path int true "Activity photo id"
//
//	@Success		200	{object}	store.ActivityPhoto
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/activityPhotos/id/{id} [get]
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

// GetActivityPhotoByTripId godoc
//
// @Summary Fetches activityPhotos by a trip id
// @Description Fetches activityPhotos by a trip id
// @Tags activityPhotos
// @Accept json
// @Produce json
// @Param id path int true "Activity id"
//
//	@Success		200	{object}	store.ActivityPhoto
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/activityPhotos/tripId/{id} [get]
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

// DeletePhotosByTripId godoc
//
// @Summary Deletes a photo
// @Description Deletes a photo by id
// @Tags activityPhotos
// @Accept json
// @Produce json
// @Param id	path		int	true	"photo ID"
//
//	@Success		204	{object} string
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/activityPhotots/id/{id} [delete]
func (app *application) deleteActivityPhotoById(w http.ResponseWriter, r *http.Request) {
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

// DeletePhotosByTripId godoc
//
// @Summary Deletes photos
// @Description Deletes photos by a trip id
// @Tags activityPhotos
// @Accept json
// @Produce json
// @Param id	path		int	true	"activity ID"
//
//	@Success		204	{object} string
//	@Failure		404	{object}	error
//	@Router			/activityPhotos/activity/{id} [delete]
//	@Failure		500	{object}	error
func (app *application) deleteActivityPhotoByActivityId(w http.ResponseWriter, r *http.Request) {
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
