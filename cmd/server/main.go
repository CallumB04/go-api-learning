package main

import (
	"fmt"
	"net/http"

	"github.com/callumb04/go-api-learning/internal/api"
)

func main() {
	// Create HTTP request multiplexer and register route handlers.
	mux := http.NewServeMux()
	api.RegisterRoutes(mux)

	// Start server on port 8080.
	fmt.Println("Server is running at localhost:8080")
	http.ListenAndServe(":8080", mux)
}
