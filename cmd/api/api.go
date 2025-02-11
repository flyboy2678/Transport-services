package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
	"transportService/docs"
	"transportService/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"go.uber.org/zap"
)

type application struct {
	config config
	store  store.Storage
	logger *zap.SugaredLogger
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)

	//Set a timeout value on the request context(ctx), that wil signal
	// throught ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	//to have nested routes it should be like this
	router.Route("/v1", func(r chi.Router) {
		docsURL := fmt.Sprintf("%s/swagger/doc.json", app.config.addr)
		r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(docsURL)))
		r.Get("/health", app.healthCheckHandler)
		//route will display as v1/health
		//users
		r.Route("/users", func(r chi.Router) {
			r.Post("/", app.createUserHandler)
			r.Route("/id/{id}", func(r chi.Router) {
				r.Get("/", app.getUserByIDHandler)
				r.Delete("/", app.deleteUserByIDHandler)
			})
			r.Route("/email/{email}", func(r chi.Router) {
				r.Get("/", app.getUserByEmailHandler)
			})
		})
		//trips
		r.Route("/trips", func(r chi.Router) {
			r.Post("/", app.createTripHandler)
			r.Get("/", app.getAllTripsHandler)
			r.Route("/id/{id}", func(r chi.Router) {
				r.Get("/", app.getTripByIdHandler)
			})
			r.Route("/location/{location}", func(r chi.Router) {
				r.Get("/", app.getTripsByLocationHandler)
			})
			r.Route("/upcoming", func(r chi.Router) {
				r.Get("/", app.getUpcomingTripsHandler)
			})
		})
		//subscriptions
		r.Route("/subscriptions", func(r chi.Router) {
			r.Post("/", app.createSubHandler)
			r.Get("/", app.getAllSubsHandler)
			r.Route("/id/{id}", func(r chi.Router) {
				r.Delete("/", app.deleteSubByUserIdHandler)
			})
			r.Route("/email/{email}", func(r chi.Router) {
				r.Delete("/", app.deleteSubByEmailHandler)
			})
		})
		//photos
		r.Route("/photos", func(r chi.Router) {
			r.Post("/", app.createPhotoHandler)
			r.Route("/id/{id}", func(r chi.Router) {
				r.Get("/", app.getPhotoByIdHandler)
				r.Delete("/", app.DeletePhotoByIdHandler)
			})
			r.Route("/tripId/{id}", func(r chi.Router) {
				r.Get("/", app.getPhotosByTripIdHandler)
				r.Delete("/", app.DeletePhotosByTripHandler)
			})
		})
	})

	router.Get("/health", app.healthCheckHandler)

	return router
}

func (app *application) run(mux http.Handler) error {
	//docs
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Version = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"

	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	shutdown := make(chan error)

	log.Printf("Server has started as %s", app.config.addr)

	err := server.ListenAndServe()

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdown
	if err != nil {
		return err
	}

	app.logger.Infow("server has stopped", "addr", app.config.addr)

	return nil
}
