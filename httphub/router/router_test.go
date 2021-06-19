package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ElMehdi19/httphub/httphub/helpers"
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

	var body structs.Response
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

			var body structs.Response
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

			var body structs.Response
			if err := json.NewDecoder(resp.Body).Decode(&body); !assert.NoError(err) {
				b, _ := ioutil.ReadAll(resp.Body)
				assert.FailNow(err.Error(), method, string(b))
			}

			assert.Equal(http.StatusOK, 200)
			assert.Equal(body.Method, method)
		})
	}
}

func TestRequest(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/request", base), nil)
	assert.NoError(err)
	req.Header.Set("user-agent", "Raymond Reddington")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(err)
	defer resp.Body.Close()

	assert.Equal(http.StatusOK, resp.StatusCode)

	var body structs.Response
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

	var body structs.Response
	err = json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(err)

	assert.Equal("127.0.0.1", body.IP)
}

func TestUserAgent(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user-agent", base), nil) //http.Get()
	assert.NoError(err)
	req.Header.Set("user-agent", "Raymond Reddington")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(err)
	defer resp.Body.Close()

	assert.Equal(http.StatusOK, resp.StatusCode)

	var body structs.Response
	err = json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(err)

	assert.Equal(req.Header.Get("user-agent"), body.UserAgent)
}

func TestHeaders(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/headers", base), nil)
	assert.NoError(err)

	headers := map[string]string{
		"Who":   "Dwight Schrute",
		"Where": "Scranton",
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(err)
	defer resp.Body.Close()

	assert.Equal(http.StatusOK, resp.StatusCode)

	var body structs.Response
	err = json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(err)

	for k, v := range headers {
		header, ok := body.Headers[k]
		assert.True(ok)
		assert.Equal(v, header)
	}
}

func TestStatusCodes(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	resp, err := http.Get(fmt.Sprintf("%s/status/xxx", base))
	assert.NoError(err)
	assert.Equal(http.StatusBadRequest, resp.StatusCode)
}

func TestBasicAuth(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	user, passwd := "mehdi", "whatever"
	resp, err := http.Get(fmt.Sprintf("%s/auth/basic/%s/%s", base, user, passwd))
	assert.NoError(err)
	assert.Equal(http.StatusUnauthorized, resp.StatusCode)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/auth/basic/%s/%s", base, user, passwd), nil)
	assert.NoError(err)
	req.SetBasicAuth(user, passwd)

	client := &http.Client{}
	resp, err = client.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
}

func TestBearerAuth(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	resp, err := http.Get(fmt.Sprintf("%s/auth/bearer", base))
	assert.NoError(err)
	assert.Equal(http.StatusUnauthorized, resp.StatusCode)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/auth/bearer", base), nil)
	assert.NoError(err)

	token := "that's what she said"
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err = client.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
}

func TestBasicAuthHidden(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	user, passwd := "mehdi", "whatever"
	resp, err := http.Get(fmt.Sprintf("%s/auth/basic-hidden/%s/%s", base, user, passwd))
	assert.NoError(err)
	assert.Equal(http.StatusNotFound, resp.StatusCode)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/auth/basic-hidden/%s/%s", base, user, passwd), nil)
	assert.NoError(err)
	req.SetBasicAuth(user, passwd)

	client := &http.Client{}
	resp, err = client.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
}

func TestResponseHeaders(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	args := map[string][]string{"X": {"a", "b"}, "Y": {"c"}}
	uri := helpers.CreateURL(base+"/response-headers", args)
	resp, err := http.Get(uri)
	assert.NoError(err)
	defer resp.Body.Close()

	var body map[string][]string
	err = json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(err)
	flat := helpers.Flatten(body)

	for key, arg := range helpers.Flatten(args) {
		val, ok := flat[key]
		assert.True(ok)
		assert.Equal(val, arg, body)
	}
}

func TestCache(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	resp, err := http.Get(fmt.Sprintf("%s/cache", base))
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
}

func TestCacheControl(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	resp, err := http.Get(fmt.Sprintf("%s/cache/1", base))
	assert.NoError(err)
	defer resp.Body.Close()
	assert.Equal(http.StatusOK, resp.StatusCode)

	var body structs.Response
	err = json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(err)
	assert.NotEmpty(body.Headers)
	assert.NotEmpty(body.URL)

	resp, err = http.Get(fmt.Sprintf("%s/cache/xx", base))
	assert.NoError(err)
	assert.Equal(http.StatusBadRequest, resp.StatusCode)
}

func TestJSONResponse(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	resp, err := http.Get(fmt.Sprintf("%s/json", base))
	assert.NoError(err)
	defer resp.Body.Close()
	assert.Equal("application/json", resp.Header.Get("content-type"))
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, resp.Body)
	assert.Equal(helpers.JSONDoc, buf.String())
}

func TestXMLResponse(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	resp, err := http.Get(fmt.Sprintf("%s/xml", base))
	assert.NoError(err)
	defer resp.Body.Close()
	assert.Equal("application/xml", resp.Header.Get("content-type"))
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, resp.Body)
	assert.Equal(helpers.XMLDoc, buf.String())
}

func TestHTMLResponse(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	resp, err := http.Get(fmt.Sprintf("%s/html", base))
	assert.NoError(err)
	defer resp.Body.Close()
	assert.Equal("text/html", resp.Header.Get("content-type"))
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, resp.Body)
	assert.Equal(helpers.HTMLDoc, buf.String())
}

func TestTXTResponse(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	resp, err := http.Get(fmt.Sprintf("%s/txt", base))
	assert.NoError(err)
	defer resp.Body.Close()
	assert.Equal("text/plain", resp.Header.Get("content-type"))
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, resp.Body)
	assert.Equal(helpers.TXTDoc, buf.String())
}

func TestCookies(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	resp, err := http.Get(fmt.Sprintf("%s/cookies", base))
	assert.NoError(err)
	defer resp.Body.Close()

	var body map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(err)
	assert.Empty(body)

	args := map[string]string{"x": "1", "y": "2"}
	uri := fmt.Sprintf("%s/cookies/set?x=%s&y=%s", base, args["x"], args["y"])
	req, err := http.NewRequest("GET", uri, nil)
	assert.NoError(err)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err = client.Do(req)
	assert.NoError(err)

	for _, cookie := range resp.Cookies() {
		val, ok := args[cookie.Name]
		assert.True(ok)
		assert.Equal(cookie.Value, val)
	}
}

func TestSetCookies(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	args := map[string][]string{
		"x": {"1"},
		"y": {"2"},
	}

	uri := helpers.CreateURL(fmt.Sprintf("%s/cookies/set", base), args)
	req, err := http.NewRequest("GET", uri, nil)
	assert.NoError(err)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	assert.NoError(err)

	assert.Len(resp.Cookies(), 2)

	cookies := helpers.Flatten(args)
	for _, c := range resp.Cookies() {
		cookie, ok := cookies[c.Name]
		assert.True(ok)
		assert.Equal(cookie, c.Value)
	}
}

func TestSetCookie(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUpTestServer()
	defer tearDown()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/cookies/set/x/1", base), nil)
	assert.NoError(err)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	assert.NoError(err)

	assert.Len(resp.Cookies(), 1)
	assert.Equal("x", resp.Cookies()[0].Name)
	assert.Equal("1", resp.Cookies()[0].Value)
}
