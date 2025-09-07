package api

import (
	"net/http"

	"github.com/callumb04/go-api-learning/internal/util"
)

// Handle access to root of api
func rootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		util.MessageResponse(w, http.StatusOK, "Server is online")
	}
}
