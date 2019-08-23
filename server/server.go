package server

import (
	"fmt"
	"net/http"
	"time"

	v1 "github.com/d3z41k/rest-boilerplate/api/v1"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// HelloWorld is a sample handler
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

// NewRouter return HTTP handler that implements the main server routers
func NewRouter() http.Handler {
	router := chi.NewRouter()

	// Set up our middleware with sane defaults
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// Set up root handler
	router.Get("/", HelloWorld)

	// Set up API
	router.Mount("/api/v1", v1.NewRouter())

	return router
}
