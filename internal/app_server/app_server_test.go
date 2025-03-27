package app_server_test

import (
	"net/http"
	"net/http/httptest"
	"recipe_app_go/internal/app_server"
	"testing"
)

func TestMuxRouter(t *testing.T) {
	router := app_server.MuxRouter()

	// Test the "/" route
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200 for '/', got %d", w.Code)
	}

	// Test the "/health" route
	req = httptest.NewRequest("GET", "/health", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200 for '/health', got %d", w.Code)
	}
}

func TestNewAppServer(t *testing.T) {
	addr := ":8080"
	server := app_server.NewAppServer(addr)

	if server.Addr != addr {
		t.Errorf("Expected server address to be %s, got %s", addr, server.Addr)
	}

	if server.Handler == nil {
		t.Error("Expected server handler to be initialized, got nil")
	}
}
