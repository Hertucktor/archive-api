package apiv1

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestStatusAlive(t *testing.T) {
	// Create a new Gin context and HTTP recorder
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call the statusAlive function
	statusAlive(c)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	expectedBody := `{"message":"OK"}`
	if w.Body.String() != expectedBody {
		t.Errorf("expected response body %q but got %q", expectedBody, w.Body.String())
	}
}

func TestStatusCheck(t *testing.T) {
	// Create a new Gin router
	router := gin.New()

	// Set up the test route
	router.GET("/status", statusCheck)

	// Create a new HTTP request to the /status route
	req, err := http.NewRequest("GET", "/status", nil)
	assert.NoError(t, err)

	// Create a new HTTP recorder to capture the response
	respRecorder := httptest.NewRecorder()

	// Send the HTTP request to the Gin router
	router.ServeHTTP(respRecorder, req)

	// Check that the response code is 200 OK
	assert.Equal(t, http.StatusOK, respRecorder.Code)

	// Check that the response body matches the expected JSON
	expected := `{"ResponseCode":200}`
	assert.Equal(t, expected, respRecorder.Body.String())
}
