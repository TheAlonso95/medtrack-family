package handler

import (
	"encoding/json"
	"net/http"

	"github.com/monorepo/backend/internal/repository"
)

// HealthResponse represents the response for the health endpoint
type HealthResponse struct {
	Status string `json:"status"`
}

// HealthHandler handles health check requests
type HealthHandler struct {
	repo repository.HealthRepository
}

// NewHealthHandler creates a new instance of HealthHandler
func NewHealthHandler(repo repository.HealthRepository) *HealthHandler {
	return &HealthHandler{
		repo: repo,
	}
}

// GetHealth handles GET requests to the /health endpoint
func (h *HealthHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	isHealthy := h.repo.CheckHealth()

	status := "ok"
	statusCode := http.StatusOK

	if !isHealthy {
		status = "error"
		statusCode = http.StatusServiceUnavailable
	}

	response := HealthResponse{
		Status: status,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}