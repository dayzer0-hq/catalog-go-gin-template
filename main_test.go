package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestHealth is the one test the starter ships with, so `go test ./...` is green from the first
// clone and you have a pattern to copy: build the router, serve a request, assert the response.
func TestHealth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	newRouter().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("GET /health = %d, want %d", w.Code, http.StatusOK)
	}
	if body := w.Body.String(); body != `{"status":"ok"}` {
		t.Fatalf("GET /health body = %q, want %q", body, `{"status":"ok"}`)
	}
}
