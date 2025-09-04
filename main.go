package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func main() {
	http.HandleFunc("/", handleRoot) // assign handler for root path ("/")

	fmt.Println("Server is running at localhost:8080")
	http.ListenAndServe(":8080", nil) // start server on port 8080
}

// define behaviour for requests to root path ("/")
func handleRoot(w http.ResponseWriter, r *http.Request) {
	// create payload for response
	response := Response{
		Message: "This is a test message.",
		Status:  200,
	}

	// set header for json content type
	w.Header().Set("Content-Type", "application/json")

	// encode to JSON and send to client
	json.NewEncoder(w).Encode(response)
}
