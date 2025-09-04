package api

import (
	"fmt"
	"net/http"

	"github.com/callumb04/playrates-backend-go/internal/models"
	"github.com/callumb04/playrates-backend-go/internal/util"
)

// Return all existing users.
func getUsers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		util.JSONResponse(w, http.StatusOK, models.TestUsers)
	})
}

// Return user that matches a provided ID.
func getUserByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get user ID from request.
		id := r.PathValue("id")

		// Search for user with matching ID
		for _, user := range models.TestUsers {
			if user.ID == id {
				util.JSONResponse(w, http.StatusOK, user)
				return
			}
		}

		// Return error if user does not exist
		util.ErrorResponse(w, http.StatusNotFound, fmt.Sprintf("User not found. ID: %s", id))
	})
}
