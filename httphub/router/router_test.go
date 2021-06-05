package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ElMehdi19/httphub/httphub/structs"
	"github.com/stretchr/testify/assert"
)

func setUpTestServer(options ...func()) (string, func()) {
	for _, opt := range options {
		opt()
	}

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
			req, err := http.NewRequest(method, fmt.Sprintf("%s/any", base), nil)
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

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user", base), nil)
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

func TestDebug(t *testing.T) {
	assert := assert.New(t)
	type testCase struct {
		name      string
		file      string
		funcOpt   func()
		isDevMode bool
		fileValid bool
	}
	tcs := []testCase{
		{
			name:    "DEV_MODE OFF",
			file:    "./router.go",
			funcOpt: func() {},
		},
		{
			name:      "DEV_MODE ON",
			file:      "./router.go",
			isDevMode: true,
			fileValid: true,
			funcOpt: func() {
				os.Setenv("DEV_MODE", "true")
			},
		},
		{
			name:      "FILE DOESNT EXIST",
			file:      "./something.go",
			isDevMode: true,
			funcOpt: func() {
				os.Setenv("DEV_MODE", "true")
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			base, tearDown := setUpTestServer(tc.funcOpt)
			defer tearDown()

			path, err := filepath.Abs(tc.file)
			assert.NoError(err)

			defer func() {
				os.Unsetenv("DEV_MODE")
			}()

			resp, err := http.Get(fmt.Sprintf("%s/debug?path=%s&line=1", base, path))
			assert.NoError(err)

			if !tc.isDevMode {
				assert.Equal(http.StatusNotFound, resp.StatusCode, os.Getenv("DEV_MODE"))
				return
			}

			if !tc.fileValid {
				assert.Equal(http.StatusInternalServerError, resp.StatusCode)
				return
			}

			assert.Equal(http.StatusOK, resp.StatusCode, os.Getenv("DEV_MODE"))
		})
	}

}

func TestIP(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	resp, err := http.Get(fmt.Sprintf("%s/ip", base))
	assert.NoError(err)
	defer resp.Body.Close()

	assert.Equal(http.StatusOK, resp.StatusCode)

	var body structs.HTTPMethodsResponse
	err = json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(err)

	assert.Equal("127.0.0.1", body.IP)
}
