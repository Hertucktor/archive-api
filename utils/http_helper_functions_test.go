package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCorsOriginMiddleware(t *testing.T) {
	// Create a mock HTTP response
	rr := httptest.NewRecorder()

	// Create a mock HTTP request with a fake URL
	req, err := http.NewRequest("GET", "http://example.com/path/to/fake/resource", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a dummy handler that simply writes a response
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err = w.Write([]byte("Hello, world!"))
		if err != nil {
			t.Error(err)
		}
	})

	// Wrap the dummy handler with the CorsOriginMiddleware
	wrappedHandler := CorsOriginMiddleware(handler)

	// Call the wrapped handler with the mock request and response
	wrappedHandler.ServeHTTP(rr, req)

	// Verify that the HTTP response has the correct Access-Control-Allow-Origin header
	if origin := rr.Header().Get("Access-Control-Allow-Origin"); origin != "*" {
		t.Errorf("handler returned wrong Access-Control-Allow-Origin header: got %v want %v", origin, "*")
	}

	// Verify that the HTTP response body contains the expected message
	expected := "Hello, world!"
	if body := rr.Body.String(); body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expected)
	}
}

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
