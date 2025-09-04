package util

import (
	"encoding/json"
	"net/http"
)

// Send JSON response to client with provided http status and data payload.
func JSONResponse(w http.ResponseWriter, status int, data any) {
	// Define content type and status code in http response header.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Encode and send http response to client
	// Function returns error, however explicitly ignoring with _
	_ = json.NewEncoder(w).Encode(data)
}

// Send Error response to client with provided http status and error message
func ErrorResponse(w http.ResponseWriter, status int, message string) {
	JSONResponse(w, status, map[string]string{"error": message})
}
