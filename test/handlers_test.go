package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"recipe_app_go/internal/handlers"
	"recipe_app_go/internal/models"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	handlers.Health(w, req)

	var response models.APIResponse
	json.NewDecoder(w.Body).Decode(&response)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
	if !response.Success || response.Data.(map[string]interface{})["status"] != "ok" {
		t.Errorf("Unexpected response: %+v", response)
	}
}
