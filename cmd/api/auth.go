package main

import (
	"fmt"
	"net/http"
	"time"
	"transportService/internal/env"
	"transportService/internal/store"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserPayload struct {
	First_name string `json:"first_name" validate:"required,max=100"`
	Last_name  string `json:"last_name" validate:"required,max=100"`
	Email      string `json:"email" validate:"required,max=255"`
	Password   string `json:"password" validate:"required,min=8,max=24"`
	Phone      string `json:"phone" validate:"required,max=255"`
}

// Register godoc
//
// @Summary Registers a user
// @Description Registers a user
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body	 RegisterUserPayload		true	"Post payload"
//
//	@Success		202		{object}	store.User
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/register [post]
func (app *application) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	hashedPassword := string(bytes)

	user := &store.User{
		First_name: payload.First_name,
		Last_name:  payload.Last_name,
		Email:      payload.Email,
		Password:   hashedPassword,
		Phone:      payload.Phone,
	}

	ctx := r.Context()

	if err := app.store.Users.Create(ctx, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	jwtToken, err := generateJWT(user)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, jwtToken); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type LogInPayload struct {
	Email    string `json:"email" validate:"required,max=255"`
	Password string `json:"password" validate:"required,min=8,max=24"`
}

// LogIn godoc
//
// @Summary Handles user log in
// @Description Handles user log in
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body	 LogInPayload		true	"Post payload"
//
//	@Success		202		{object}	string
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/login [post]
func (app *application) logInHandler(w http.ResponseWriter, r *http.Request) {
	var payload LogInPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	user, err := app.store.Users.GetByEmail(ctx, payload.Email)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	fmt.Print("User Password", user.Password)
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password), []byte(payload.Password),
	); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	jwtToken, err := generateJWT(user)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, jwtToken); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func generateJWT(user *store.User) (string, error) {
	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	key = []byte(env.GetString("JWT_SECRET", "secret"))

	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":         user.ID,
			"first_name": user.First_name,
			"last_name":  user.Last_name,
			"email":      user.Email,
			"Phone":      user.Phone,
			"exp":        time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return s, nil
}
