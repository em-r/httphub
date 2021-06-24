package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViewRedirect(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000", nil)
	assert.NoError(err)

	rec := httptest.NewRecorder()
	viewRedirect(rec, req, "http://httphub.io")

	assert.Equal(http.StatusFound, rec.Result().StatusCode)
	assert.Equal("http://httphub.io", rec.Result().Header.Get("Location"))
}
