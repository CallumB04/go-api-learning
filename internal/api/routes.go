package api

import (
	"net/http"
)

// Register route handlers to mux.
func RegisterRoutes(mux *http.ServeMux) {
	// User Routes
	mux.Handle("GET /users", getUsers())
}
