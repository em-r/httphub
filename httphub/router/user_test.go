package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ElMehdi19/httphub/httphub/helpers"
	"github.com/ElMehdi19/httphub/httphub/structs"
	"github.com/stretchr/testify/assert"
)

func TestViewUser(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000", nil)
	assert.NoError(err)

	req.RemoteAddr = "127.0.0.1"
	headers := map[string]string{
		"user-agent": "Raymond Reddington",
		"passcode":   "whatever",
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	rec := httptest.NewRecorder()
	ViewUser(rec, req)

	var body structs.Response
	err = json.NewDecoder(rec.Body).Decode(&body)
	assert.NoError(err)

	assert.Equal(req.RemoteAddr, body.IP)
	assert.Equal(req.Header.Get("user-agent"), body.UserAgent)
	assert.Equal(fmt.Sprintf("%v", helpers.Flatten(req.Header)), fmt.Sprintf("%v", body.Headers))
}

func TestViewIP(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/ip", nil)
	assert.NoError(err)
	req.RemoteAddr = "127.0.0.1"

	rec := httptest.NewRecorder()
	ViewIP(rec, req)

	var body structs.Response
	err = json.NewDecoder(rec.Body).Decode(&body)
	assert.NoError(err)

	assert.Equal(req.RemoteAddr, body.IP)
}

func TestViewUserAgent(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000", nil)
	assert.NoError(err)
	req.Header.Set("user-agent", "Raymond Reddington")

	rec := httptest.NewRecorder()
	ViewUserAgent(rec, req)

	var body structs.Response
	err = json.NewDecoder(rec.Body).Decode(&body)
	assert.NoError(err)

	assert.Equal(req.Header.Get("user-agent"), body.UserAgent)
}
