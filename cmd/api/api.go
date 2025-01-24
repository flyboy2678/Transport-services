package main

import (
	"log"
	"net/http"
	"time"
	"transportService/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store store.Storage
}

type config struct {
	addr string
	db dbConfig
}

type dbConfig struct{
	addr string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime string 
}

func (app* application) mount() http.Handler{
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	
	//Set a timeout value on the request context(ctx), that wil signal
	// throught ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60* time.Second))
	
	//to have nested routes it should be like this
	router.Route("/v1", func(r chi.Router){
		r.Get("/health", app.healthCheckHandler)
		//route will display as v1/health
	})
	
	router.Get("/health", app.healthCheckHandler)

	return router
}

func (app *application) run(mux http.Handler) error{
	server := &http.Server{
		Addr: app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("Server has started as %s", app.config.addr)

	return server.ListenAndServe()
}