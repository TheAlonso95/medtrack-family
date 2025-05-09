package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockHealthRepository is a mock implementation of the HealthRepository interface
type MockHealthRepository struct {
	healthy bool
}

func (m *MockHealthRepository) CheckHealth() bool {
	return m.healthy
}

func TestGetHealth(t *testing.T) {
	tests := []struct {
		name           string
		isHealthy      bool
		expectedStatus string
		expectedCode   int
	}{
		{
			name:           "healthy system",
			isHealthy:      true,
			expectedStatus: "ok",
			expectedCode:   http.StatusOK,
		},
		{
			name:           "unhealthy system",
			isHealthy:      false,
			expectedStatus: "error",
			expectedCode:   http.StatusServiceUnavailable,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock repository
			mockRepo := &MockHealthRepository{healthy: tt.isHealthy}
			
			// Create the handler with the mock repository
			handler := NewHealthHandler(mockRepo)
			
			// Create a request to pass to the handler
			req := httptest.NewRequest("GET", "/health", nil)
			
			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()
			
			// Call the handler
			handler.GetHealth(rr, req)
			
			// Check the status code
			if status := rr.Code; status != tt.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedCode)
			}
			
			// Check the response body
			var response HealthResponse
			err := json.NewDecoder(rr.Body).Decode(&response)
			if err != nil {
				t.Errorf("failed to decode response body: %v", err)
			}
			
			if response.Status != tt.expectedStatus {
				t.Errorf("handler returned unexpected status: got %v want %v", response.Status, tt.expectedStatus)
			}
		})
	}
}