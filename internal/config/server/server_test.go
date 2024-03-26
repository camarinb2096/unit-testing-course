package server

import "testing"

// TestNewServer is a test for NewServer function
func TestNewServer(t *testing.T) {
	t.Run("NewServer", func(t *testing.T) {
		// NewServer is a function that returns a new Server instance
		got := NewServer()
		// We expect the Router field to be initialized
		if got.router == nil {
			t.Errorf("NewServer().Router = nil; want initialized")
		}
	})
}

// TestRoutes is a test for Routes function
func TestRoutes(t *testing.T) {
	t.Run("Routes", func(t *testing.T) {
		// NewServer is a function that returns a new Server instance
		s := NewServer()
		// Routes is a function that initializes the routes
		s.Routes()
		// We expect the Router field to be initialized
		if s.router == nil {
			t.Errorf("Routes().Router = nil; want initialized")
		}
	})
}
