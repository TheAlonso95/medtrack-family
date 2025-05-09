package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/monorepo/backend/internal/repository"
)

func TestHealthEndpoint(t *testing.T) {
	// Create a new router
	r := chi.NewRouter()

	// Create a real repository
	healthRepo := repository.NewHealthRepository()

	// Create a handler with the repository
	healthHandler := NewHealthHandler(healthRepo)

	// Register the health endpoint
	r.Get("/health", healthHandler.GetHealth)

	// Create a request to the health endpoint
	req := httptest.NewRequest("GET", "/health", nil)
	
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	
	// Serve the request
	r.ServeHTTP(rr, req)
	
	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	
	// Check the response body
	var response HealthResponse
	err := json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}
	
	if response.Status != "ok" {
		t.Errorf("handler returned unexpected status: got %v want %v", response.Status, "ok")
	}
}