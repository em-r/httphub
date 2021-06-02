package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ElMehdi19/httphub/httphub/structs"
	"github.com/stretchr/testify/assert"
)

func TestViewUser(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000", nil)
	assert.NoError(err)

	req.RemoteAddr = "127.0.0.1"
	req.Header.Set("user-agent", "Raymond Reddington")

	rec := httptest.NewRecorder()
	ViewUser(rec, req)

	var body structs.HTTPMethodsResponse
	err = json.NewDecoder(rec.Body).Decode(&body)
	assert.NoError(err)

	assert.Equal(req.RemoteAddr, body.IP)
	assert.Equal(req.Header.Get("user-agent"), body.UserAgent)
}
