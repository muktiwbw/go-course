package main

import (
	"log"

	"github.com/muktiwbw/go-course/scripts/env"
)

// Mainly used for setting up configurations and dependencies
/**
 * TODO 7 (DONE)
 * Create an env util that returns the value of given key
 */
func main() {
	cfg := config{
		addr: env.GetString("APP_ADDRESS", ":3000"),
	}

	// Initialize application from the api.go
	app := &application{
		config: cfg,
	}

	/**
	 * TODO 2 (DONE)
	 * Call the run method from app
	 * Log in case there's error using log.Fatal(app.run())
	 */
	mux := app.mount()
	log.Fatal(app.run(mux))
}
