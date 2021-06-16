package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ElMehdi19/httphub/httphub/structs"
	"github.com/stretchr/testify/assert"
)

func TestViewCookies(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000", nil)
	assert.NoError(err)
	rec := httptest.NewRecorder()
	ViewCookies(rec, req)

	var body structs.Response
	err = json.NewDecoder(rec.Body).Decode(&body)
	assert.NoError(err)
	assert.Empty(body.Cookies)
}
