package main

import (
	"net/http"
	"strconv"
	"transportService/internal/store"

	"github.com/go-chi/chi/v5"
)

type CreateCommentPayload struct {
	User_id int64  `json:"user_id" validate:"required"`
	Trip_id int64  `json:"trip_id" validate:"required"`
	Comment string `json:"comment" validate:"required"`
	Rating  int    `json:"rating" validate:"required"`
}

// CreateComment godoc
//
// @Summary Creates a comment
// @Description Creates a comment
// @Tags comments
// @Accept json
// @Produce json
// @Param payload body	 CreateCommentPayload		true	"Post payload"
//
//	@Success		202		{object}	store.Comment
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/comments [post]
func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateCommentPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	comment := &store.Comment{
		User_id: payload.User_id,
		Trip_id: payload.Trip_id,
		Comment: payload.Comment,
		Rating:  payload.Rating,
	}

	ctx := r.Context()

	if err := app.store.Comments.Create(ctx, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

}

// GetCommentsById godoc
//
// @Summary Fetches comments by id
// @Description Fetches comments by id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "comment id"
//
//	@Success		200	{object}	store.Comment
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/comments/id/{id} [get]
func (app *application) getCommentByIdHandler(w http.ResponseWriter, r *http.Request) {
	commentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	comment, err := app.store.Comments.GetByID(ctx, commentId)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetCommentsByTripId godoc
//
// @Summary Fetches comments by a trip id
// @Description Fetches comments by a trip id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Trip id"
//
//	@Success		200	{object}	store.Comment
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/comments/tripId/{id} [get]
func (app *application) getCommentsByTripIdHandler(w http.ResponseWriter, r *http.Request) {
	tripId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	comments, err := app.store.Comments.GetByTripID(ctx, tripId)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, comments); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// DeleteCommentsById godoc
//
// @Summary Deletes a comment
// @Description Deletes a comment by id
// @Tags comments
// @Accept json
// @Produce json
// @Param id	path		int	true	"Comment ID"
//
//	@Success		204	{object} string
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/comments/id/{id} [delete]
func (app *application) deleteCommentByIdHandler(w http.ResponseWriter, r *http.Request) {
	commentId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Comments.DeleteByID(ctx, commentId); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteCommentsByTripId godoc
//
// @Summary Deletes comments
// @Description Deletes comments by a trip id
// @Tags comments
// @Accept json
// @Produce json
// @Param id	path		int	true	"trip ID"
//
//	@Success		204	{object} string
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/comments/tripId/{id} [delete]
func (app *application) deleteCommentByTripIdHandler(w http.ResponseWriter, r *http.Request) {
	tripId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Comments.DeleteByTripID(ctx, tripId); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
