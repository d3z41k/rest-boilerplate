package v1

import (
	"net/http"

	"github.com/d3z41k/rest-boilerplate/controllers"
	m "github.com/d3z41k/rest-boilerplate/middleware"

	"github.com/go-chi/chi"
)

// NewRouter returns an HTTP handler that implements the routes for the API
func NewRouter() http.Handler {
	r := chi.NewRouter()

	// Attach JWT auth middleware
	r.Use(m.JwtAuthentication) 

	// Register the API routes
	r.Post("/api/user/new", controllers.CreateAccount)
	r.Post("/api/user/login", controllers.Authenticate)
	r.Post("/api/contacts/new", controllers.CreateContact)
	r.Get("/api/contacts", controllers.GetContactsFor)

	return r
}
