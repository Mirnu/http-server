package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHello(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, "Hello", rec.Body.String())
}
