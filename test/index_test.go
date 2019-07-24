package test

import (
	"github.com/stretchr/testify/assert"
	"my-blog/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexRoute(t *testing.T) {
	app := router.Init()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	app.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "The status code should be equal")
	assert.Equal(t, "Hello World!", w.Body.String(), "The response text should be equal")
}
