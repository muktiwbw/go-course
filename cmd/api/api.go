package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type config struct {
	addr string
}

// This file is where the app will live
type application struct {
	config config
}

/**
 * TODO 4 (DONE)
 * * Mux, which is a Handler is very similar to Route if you've coded in other language like Express or Laravel
 * * This is where you list all of your routes (or handler in Go)
 * Separate Mux config to another method
 * - mount()
 * - return *http.ServeMux()
 */
/**
 * TODO 5 (DONE)
 * ? Create a separate package containing separate files for each handler category
 * ! Cant' do that since it wouldn't know the existence of application struct
 * Create a handler for server health check.
 * - /healthcheck
 * - return "ok"
 * - Add mux.HandleFunc to add new route
 */
/**
 * TODO 6 (DONE)
 * Add the api version in the handler group using chi library
 * - Add v1 handler group
 * - Add logger middleware from chi
 * - Add recover (handles recovery when panic) middleware from chi
 * - Update mount() to return http.Handler instead
 *
 * * The reason why we want to return http.Handler is to keep abstraction in the run() method
 * * If in the future we want to replace our handler with, let's say Gin
 * * we don't need to change the implementation on the run() method.
 * * We only pass the signature argument, which is http.Handler which they (chi.Mux and gin.Handler(probably)) implement
 */
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

/**
 * TODO 1 (DONE)
 * Add a method in the application struct
 * - run()
 * - returns error
 * - create http.Server
 * - return ListenAndServe
 * - contains Mux handler -> http.NewServeMux
 * - add logger when running server
 */
/**
 * TODO 3 (DONE)
 * Add timeout for both read and write
 * - write timeout: 30
 * - read timeout: 10
 * - idle timeout: 1 min
 */
func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server is running on %s", app.config.addr)

	return srv.ListenAndServe()
}
