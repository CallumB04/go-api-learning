package api

import (
	"fmt"
	"net/http"

	"github.com/callumb04/go-api-learning/internal/data"
	"github.com/callumb04/go-api-learning/internal/util"
)

// Return all existing users.
func getUsers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Load users and return JSON to client
		if users, err := data.LoadUsers(); err == nil {
			util.JSONResponse(w, http.StatusOK, users)
			return
		}

		// Return error to client if users fetch fails
		util.ErrorResponse(w, http.StatusInternalServerError, "Error fetching users")

	})
}

// Return user that matches a provided ID.
func getUserByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get user ID from request.
		id := r.PathValue("id")

		// Load users to search with ID
		if users, err := data.LoadUsers(); err == nil {

			// Search for user with matching ID
			for _, user := range users {
				if user.ID == id {
					util.JSONResponse(w, http.StatusOK, user)
					return
				}
			}

			// Return error if user does not exist
			util.ErrorResponse(w, http.StatusNotFound, fmt.Sprintf("User not found. ID: %s", id))
			return
		}
		// Return error to client if users fetch fails
		util.ErrorResponse(w, http.StatusInternalServerError, "Error fetching users")
	})
}
