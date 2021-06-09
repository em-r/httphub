package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViewCache(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/cache", nil)
	assert.NoError(err)

	rec := httptest.NewRecorder()
	ViewCache(rec, req)

	assert.Equal(http.StatusOK, rec.Result().StatusCode)
	assert.NotEmpty(rec.Header().Get("Date"))
	assert.NotEmpty(rec.Header().Get("Etag"))
	assert.NotEmpty(rec.Header().Get("Content-Location"))

	req.Header.Set("If-None-Match", "something-random")
	rec = httptest.NewRecorder()
	ViewCache(rec, req)
	assert.Equal(http.StatusNotModified, rec.Result().StatusCode)
}
