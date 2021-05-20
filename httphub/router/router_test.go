package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/ElMehdi19/httphub/httphub/structs"
	"github.com/stretchr/testify/assert"
)

func setUpTestServer() (string, func()) {
	mux := New()
	server := httptest.NewServer(mux)
	return server.URL, func() {
		server.Close()
	}
}

func TestMethodGet(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	resp, err := http.Get(fmt.Sprintf("%s/get?x=1", base))
	assert.NoError(err)
	assert.Equal(resp.StatusCode, http.StatusOK)

	var body structs.HTTPMethodsResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); !assert.NoError(err) {
		assert.FailNow(err.Error())
	}
	assert.True(reflect.DeepEqual(map[string][]string{"x": {"1"}}, body.Args))
}

func TestMethodPost(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	resp, err := http.Post(fmt.Sprintf("%s/post?x=1", base), "text/plain", strings.NewReader("test"))
	assert.NoError(err)
	assert.Equal(resp.StatusCode, http.StatusOK)

	var body structs.HTTPMethodsResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); !assert.NoError(err) {
		assert.FailNow(err.Error())
	}

	assert.True(reflect.DeepEqual(map[string][]string{"x": {"1"}}, body.Args))
	assert.Equal("test", body.Data)
	assert.Equal(body.Headers["Content-Type"][0], "text/plain")
}
