package server

import (
	"fmt"
	"net/http"
	"time"

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
	router.Use(middleware.DefaultCompress)
	router.Use(middleware.Timeout(60 * time.Second))

	//Set up root handler
	router.Get("/", HelloWorld)

	return router
}
