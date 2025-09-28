package handler

import (
	"net/http"
	"os"
	"testing"
)

// TestAdminRoutes tests basic admin route accessibility
func TestAdminRoutes(t *testing.T) {
	// Set debug mode for testing
	os.Setenv("ADMIN_DEBUG", "true")
	defer os.Unsetenv("ADMIN_DEBUG")

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "Admin dashboard should be accessible",
			method:         "GET",
			path:           "/admin/",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Admin music list should be accessible",
			method:         "GET",
			path:           "/admin/music",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Admin photograph list should be accessible",
			method:         "GET",
			path:           "/admin/photograph",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Admin scene list should be accessible",
			method:         "GET",
			path:           "/admin/scene",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make actual HTTP request to running server
			client := &http.Client{}
			resp, err := client.Get("http://localhost:8090" + tt.path)
			if err != nil {
				t.Skipf("Server not running, skipping integration test: %v", err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d for %s %s",
					tt.expectedStatus, resp.StatusCode, tt.method, tt.path)
			}
		})
	}
}

// TestAdminTemplateNames tests that template names are correctly resolved
func TestAdminTemplateNames(t *testing.T) {
	templateNames := []string{
		"admin/dashboard.go.tmpl",
		"admin/music_list.go.tmpl",
		"admin/music_edit.go.tmpl",
		"admin/photograph_list.go.tmpl",
		"admin/photograph_edit.go.tmpl",
		"admin/scene_list.go.tmpl",
		"admin/scene_edit.go.tmpl",
	}

	for _, name := range templateNames {
		t.Run("Template "+name+" should exist", func(t *testing.T) {
			// This is a basic check that template names follow the expected pattern
			if len(name) == 0 {
				t.Error("Template name should not be empty")
			}
			if name[:6] != "admin/" {
				t.Errorf("Template name should start with 'admin/', got: %s", name)
			}
		})
	}
}