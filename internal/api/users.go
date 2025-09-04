package api

import (
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
