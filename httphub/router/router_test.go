package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
	assert.Equal("application/json", resp.Header.Get("content-type"))

	var body structs.HTTPMethodsResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); !assert.NoError(err) {
		assert.FailNow(err.Error())
	}
	assert.Equal(fmt.Sprintf("%v", map[string]string{"x": "1"}), fmt.Sprintf("%v", body.Args))

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
			defer resp.Body.Close()

			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.Equal("application/json", resp.Header.Get("content-type"))

			var body structs.HTTPMethodsResponse
			if err := json.NewDecoder(resp.Body).Decode(&body); !assert.NoError(err) {
				b, _ := ioutil.ReadAll(resp.Body)
				assert.FailNow(err.Error(), string(b))
			}

			assert.Equal(fmt.Sprintf("%v", map[string]string{"x": "1"}), fmt.Sprintf("%v", body.Args))
			assert.Equal("test", body.Data)
			assert.Equal(body.Headers["Content-Type"], "text/plain")
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

			assert.Equal("application/json", resp.Header.Get("content-type"))

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

func TestUser(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", base, "user"), nil)
	assert.NoError(err)
	req.Header.Set("user-agent", "Raymond Reddington")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(err)
	defer resp.Body.Close()

	assert.Equal(http.StatusOK, resp.StatusCode)

	var body structs.HTTPMethodsResponse
	err = json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(err)

	assert.Equal(req.Header.Get("user-agent"), body.UserAgent)
	assert.Equal("127.0.0.1", body.IP)
}
