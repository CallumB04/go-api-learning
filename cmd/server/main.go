package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define structure for JSON response.
type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func main() {
	// Assign handler for root path.
	http.HandleFunc("/", handleRoot)

	fmt.Println("Server is running at localhost:8080")

	// Start server on port 8080.
	http.ListenAndServe(":8080", nil)
}

// Handle requests to the root path.
func handleRoot(w http.ResponseWriter, r *http.Request) {
	// Create payload for response.
	response := Response{
		Message: "This is a test message.",
		Status:  200,
	}

	// Set header for json content type.
	w.Header().Set("Content-Type", "application/json")

	// Encode response to JSON and send to client.
	json.NewEncoder(w).Encode(response)
}
