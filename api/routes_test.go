package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLoginRoute(t *testing.T) {
	router := InitRoute()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
