package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetUserByID(t *testing.T) {

	tests := []struct {
		name           string
		userID         string
		expectedStatus int
		expectedBody   string
	}{
		// TODO: Implement other tests, currently can't load users from file
		// {
		// 	"valid user", "1", http.StatusOK,
		// 	`{
		// 		"id": "1",
		// 		"email": "thisisanemail@gmail.com",
		// 		"username": "callum",
		// 		"bio": "This is a test bio",
		// 		"createdAt": "2025-09-05T12:34:56Z"
		// 	}`,
		// },
		// {"user not found", "9999", http.StatusNotFound, `{"error":"User not found"}`},
		{"missing user id", "", http.StatusBadRequest, `{"error":"Missing user ID"}`},
	}

	mux := http.NewServeMux()
	RegisterRoutes(mux)

	for _, tt := range tests {
		// Create http request to send to handler.
		req := httptest.NewRequest(http.MethodGet, "/users/"+tt.userID, nil)
		w := httptest.NewRecorder()

		// Send http request to handler.
		mux.ServeHTTP(w, req)

		// Receive response from the handler function.
		resp := w.Result()
		defer resp.Body.Close() // close response body after test finishes.

		// Check response sends correct status code.
		if resp.StatusCode != tt.expectedStatus {
			t.Errorf("Expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
		}

		// Parse JSON body from response.
		var body map[string]any
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			t.Errorf("Failed to decode response: %v", err)
		}

		// Check body (message/error/user) matches result.
		rawBody := strings.TrimSpace(w.Body.String())
		if rawBody != tt.expectedBody {
			t.Errorf("Expected body %q, got %q", tt.expectedBody, rawBody)
		}
	}

}
