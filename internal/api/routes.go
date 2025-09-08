package api

import (
	"net/http"
)

// Register route handlers to mux.
func RegisterRoutes(mux *http.ServeMux) {
	// Root
	mux.Handle("GET /", rootHandler())
	// User Routes
	mux.Handle("GET /users", getUsers())
	mux.Handle("GET /users/", getUserByID())
	mux.Handle("GET /users/{id}", getUserByID())
}
