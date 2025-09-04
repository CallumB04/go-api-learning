package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRoot) // assign handler for root path ("/")

	fmt.Println("Server is running at localhost:8080")
	http.ListenAndServe(":8080", nil) // start server on port 8080
}

// define behaviour for requests to root path ("/")
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
