package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestLogger(t *testing.T) {
	// Setup Gin with Logger middleware
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(Logger())

	// Dummy handler to simulate a request
	r.GET("/test", func(c *gin.Context) {
		time.Sleep(10 * time.Millisecond) // Simulate some processing time
		c.String(http.StatusOK, "test response")
	})

	// Capture the log output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(nil) // Reset the output to default
	}()

	// Create a test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)

	// Perform the request
	r.ServeHTTP(w, req)

	// Test the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check if the log contains expected information
	logOutput := buf.String()
	if !bytes.Contains([]byte(logOutput), []byte("GET")) || !bytes.Contains([]byte(logOutput), []byte("/test")) {
		t.Errorf("Log output does not contain expected method or path")
	}

	// Here you could add more checks, for instance, to verify the IP and the duration format
}
