package router

import (
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestNew(t *testing.T) {
	r := New()

	// Check if the router is a chi.Mux
	if _, ok := interface{}(r).(*chi.Mux); !ok {
		t.Error("New() did not return a *chi.Mux")
	}

	// Check if middleware is configured
	// This is a simple check to ensure the router has middleware
	// In a real test, you might want to check specific middleware
	middlewareCount := len(r.Middlewares())
	if middlewareCount == 0 {
		t.Error("Router has no middleware configured")
	}
}