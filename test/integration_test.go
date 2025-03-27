package integration_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"recipe_app_go/internal/app_server"
	"recipe_app_go/internal/models"
	"testing"
)

// TestMain sets up the server for all integration tests.
func TestMain(m *testing.M) {
	// Setup code before running tests (e.g., init database)
	// teardown() can be deferred if necessary

	m.Run() // Run the actual tests
}

// TestHealthCheckIntegration verifies the API is up and responding.
func TestHealthCheckIntegration(t *testing.T) {
	// Spin up a test server
	ts := httptest.NewServer(app_server.MuxRouter()) // Assuming NewRouter() sets up routes
	defer ts.Close()

	// Make request to test server
	resp, err := http.Get(ts.URL + "/health")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Validate response status
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// Decode JSON response
	var response models.APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	// Validate response structure
	if !response.Success || response.Data.(map[string]interface{})["status"] != "ok" {
		t.Errorf("Unexpected response: %+v", response)
	}
}

// TestHomeIntegration verifies the home route.
func TestHomeIntegration(t *testing.T) {
	ts := httptest.NewServer(app_server.MuxRouter())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var response models.APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	if !response.Success || response.Data.(map[string]interface{})["message"] != "hi" {
		t.Errorf("Unexpected response: %+v", response)
	}
}
