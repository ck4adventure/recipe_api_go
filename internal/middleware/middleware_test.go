package middleware_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"recipe_app_go/internal/middleware"
	"strings"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	// Create a test handler to be wrapped by the middleware
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Wrap the test handler with the LoggingMiddleware
	handler := middleware.LoggingMiddleware(testHandler)

	// Create a test request and response recorder
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Capture the log output
	var logOutput strings.Builder
	log.SetOutput(&logOutput)

	// Serve the request
	handler.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Check the log output
	expectedLog := "request GET /"
	if !strings.Contains(logOutput.String(), expectedLog) {
		t.Errorf("Expected log output to contain %q, got %q", expectedLog, logOutput.String())
	}
}
