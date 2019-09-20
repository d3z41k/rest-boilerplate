package v1

import (
	"net/http"

	"github.com/d3z41k/rest-boilerplate/controllers"
	m "github.com/d3z41k/rest-boilerplate/middleware"

	"github.com/go-chi/chi"
)

// NewRouter returns an HTTP handler that implements the routes for the API
func NewRouter() http.Handler {
	router := chi.NewRouter()

	// Attach JWT auth middleware
	router.Use(m.JwtAuthentication)

	// Register the API routes
	router.Post("/user/new", controllers.CreateAccount)
	router.Post("/user/login", controllers.Authenticate)
	router.Post("/contact/new", controllers.CreateContact)
	router.Get("/contacts", controllers.GetContacts)
	router.Get("/contact/{id}", controllers.GetContact)
	router.Put("/contact/{id}", controllers.UpdateContact)
	router.Delete("/contact/{id}", controllers.DeleteContact)

	return router
}
