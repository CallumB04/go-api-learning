package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/callumb04/go-api-learning/internal/data"
	"github.com/callumb04/go-api-learning/internal/util"
)

// Return all existing users.
func getUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Load users from local JSON file.
		users, err := data.LoadUsers()

		// Return error to client if users fetch fails
		if err != nil {
			util.ErrorResponse(w, http.StatusInternalServerError, "Error fetching users")
			return
		}

		// Send users to client.
		util.JSONResponse(w, http.StatusOK, users)

	}
}

// Return user that matches a provided ID.
func getUserByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get user ID from request and remove whitespace
		id := strings.TrimSpace(r.PathValue("id"))

		// Return error if missing user ID.
		if id == "" {
			util.ErrorResponse(w, http.StatusBadRequest, "Missing user ID")
			return
		}

		users, err := data.LoadUsers()

		if err != nil {
			// Return error to client if users fetch fails.
			util.ErrorResponse(w, http.StatusInternalServerError, "Error fetching users")
			return
		}

		// Search for user with matching ID and return to client.
		for _, user := range users {
			if user.ID == id {
				util.JSONResponse(w, http.StatusOK, user)
				return
			}
		}

		// Return error if user does not exist.
		util.ErrorResponse(w, http.StatusNotFound, fmt.Sprintf("User not found. ID: %s", id))

	}
}
