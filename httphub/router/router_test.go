package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func TestGET(t *testing.T) {
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

	resp, _ = http.Post(fmt.Sprintf("%s/get?x=1", base), "", nil)
	assert.Equal(http.StatusMethodNotAllowed, resp.StatusCode)
}

func TestMethods(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	tcs := []string{"post", "put", "patch", "delete"}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			method := strings.ToUpper(tc)
			req, err := http.NewRequest(method, fmt.Sprintf("%s/%s?x=1", base, tc), strings.NewReader("test"))
			assert.NoError(err)
			req.Header.Set("content-type", "text/plain")

			client := &http.Client{}
			resp, err := client.Do(req)
			assert.NoError(err)

			assert.Equal(resp.StatusCode, http.StatusOK)

			var body structs.HTTPMethodsResponse
			if err := json.NewDecoder(resp.Body).Decode(&body); !assert.NoError(err) {
				b, _ := ioutil.ReadAll(resp.Body)
				assert.FailNow(err.Error(), string(b))
			}

			assert.True(reflect.DeepEqual(map[string][]string{"x": {"1"}}, body.Args))
			assert.Equal("test", body.Data)
			assert.Equal(body.Headers["Content-Type"][0], "text/plain")
		})
	}
}

func TestAny(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	methods := []string{"GET", "PUT", "POST", "PUT", "DELETE"}
	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", base, "any"), nil)
			assert.NoError(err)

			client := &http.Client{}
			resp, err := client.Do(req)
			assert.NoError(err)
			defer resp.Body.Close()

			var body structs.HTTPMethodsResponse
			if err := json.NewDecoder(resp.Body).Decode(&body); !assert.NoError(err) {
				b, _ := ioutil.ReadAll(resp.Body)
				assert.FailNow(err.Error(), method, string(b))
			}

			assert.Equal(http.StatusOK, 200)
			assert.Equal(body.Method, method)
		})
	}
}
