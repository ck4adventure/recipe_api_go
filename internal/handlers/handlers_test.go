package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"recipe_app_go/internal/handlers"
	"recipe_app_go/internal/models"
	"testing"
)

// Test HealthCheck endpoint
func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.Health)
	handler.ServeHTTP(rr, req)

	// Validate status code
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v",
			rr.Code, http.StatusOK)
	}

	// Validate JSON response
	var response models.APIResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode JSON response: %v", err)
	}

	// Ensure success flag is true
	if !response.Success {
		t.Errorf("Expected Success=true, got %v", response.Success)
	}

	// Ensure the response contains the expected "status":"ok" data
	data, ok := response.Data.(map[string]interface{})
	if !ok || data["status"] != "ok" {
		t.Errorf("Unexpected data response: %+v", response.Data)
	}
}

// Test Home endpoint
func TestHome(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handlers.Home(w, req)

	// Decode response
	var response models.APIResponse
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode JSON response: %v", err)
	}

	// Validate status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	// Ensure success flag is true
	if !response.Success {
		t.Errorf("Expected Success=true, got %v", response.Success)
	}

	// Ensure the response contains the expected "message":"hi" data
	data, ok := response.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected Data to be a map, got %T", response.Data)
	}

	if message, exists := data["message"]; !exists || message != "hi" {
		t.Errorf("Unexpected message in response: %+v", data)
	}
}
