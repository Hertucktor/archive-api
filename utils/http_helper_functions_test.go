package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotFoundHandler(t *testing.T) {
	// Create a new Gin router and register the NotFoundHandler
	router := gin.Default()
	router.NoRoute(NotFoundHandler)

	// Create a new HTTP request to trigger the NotFoundHandler
	req, err := http.NewRequest("GET", "/path/to/nonexistent/resource", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Use httptest.NewRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()

	// Call the router's ServeHTTP method to process the request and response
	router.ServeHTTP(recorder, req)

	// Assert that the response status code is 404
	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, recorder.Code)
	}

	// Assert that the response body contains the expected message
	expectedBody := `{"message":"Page not found"}`
	if recorder.Body.String() != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, recorder.Body.String())
	}
}
