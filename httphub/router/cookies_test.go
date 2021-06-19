package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

func TestViewSetCookies(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/", nil)
	assert.NoError(err)

	args := map[string]string{
		"x": "1",
		"y": "2",
	}
	req.URL.RawQuery = "x=1&y=2"

	rec := httptest.NewRecorder()
	ViewSetCookies(rec, req)
	assert.Equal(http.StatusFound, rec.Result().StatusCode)
	assert.Len(rec.Result().Cookies(), 2, fmt.Sprintf("args: %s", req.URL.Query()))

	for _, c := range rec.Result().Cookies() {
		cookie, ok := args[c.Name]
		assert.True(ok)
		assert.Equal(cookie, c.Value)
	}
}

func TestViewSetCookie(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000", nil)
	assert.NoError(err)

	rec := httptest.NewRecorder()
	ViewSetCookie(rec, req)

	assert.Equal(http.StatusFound, rec.Result().StatusCode)
}

func TestViewDeleteCookies(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/set?names=x", nil)
	assert.NoError(err)
	req.AddCookie(&http.Cookie{
		Name:  "x",
		Value: "1",
	})

	rec := httptest.NewRecorder()
	ViewDeleteCookies(rec, req)

	assert.Equal(http.StatusFound, rec.Result().StatusCode)

	cookies := rec.Result().Cookies()
	assert.Len(cookies, 1)

	c := cookies[0]
	assert.Equal("x", c.Name)
	assert.True(c.Expires.Before(time.Now()))
}
