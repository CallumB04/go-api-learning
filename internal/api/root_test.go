package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	// Create http request to send to handler.
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder() // http response writer

	// Send http request to handler.
	handler := rootHandler()
	handler(w, req)

	// Receive response from the handler function.
	resp := w.Result()
	defer resp.Body.Close() // close response body after test finishes.

	// Check response contains correct status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// Parse JSON body from response.
	var body map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Errorf("Failed to decode response: %v", err)
	}

	// Check message in body matches expected message
	expectedMessage := "Server is online"
	if body["message"] != expectedMessage {
		t.Errorf("Expected message %q, got %q", expectedMessage, body["message"])
	}
}
