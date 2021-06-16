package helpers

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/ElMehdi19/httphub/httphub/structs"
	"github.com/stretchr/testify/assert"
)

func TestParseBody(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000", nil)
	assert.NoError(err)

	var resp structs.Response
	parseBody(req, &resp)
	assert.Empty(resp.JSON)
	assert.Empty(resp.Form)
	assert.Empty(resp.Data)

	type testCase struct {
		contentType, body string
	}

	tcs := []testCase{
		{
			contentType: "application/x-www-form-urlencoded",
			body:        "x=1&y=2",
		},
		{
			contentType: "text/plain",
			body:        "whatever",
		},
		{
			contentType: "application/json",
			body: `{
				"something": "random"
			}`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.contentType, func(t *testing.T) {
			req, err = http.NewRequest("POST", "http://127.0.0.1:5000", strings.NewReader(tc.body))
			req.Header.Set("content-type", tc.contentType)
			var resp structs.Response
			parseBody(req, &resp)
			switch tc.contentType {
			case "application/json":
				assert.Empty(resp.Data)
				assert.Empty(resp.Form)
				assert.NotEmpty(resp.JSON)
			case "application/x-www-form-urlencoded":
				assert.Empty(resp.Data)
				assert.Empty(resp.JSON)
				assert.NotEmpty(resp.Form)
			case "text/plain":
				assert.Empty(resp.JSON)
				assert.Empty(resp.Form)
				assert.NotEmpty(resp.Data)
			}
		})
	}

}

func TestMakeResponse(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000", nil)
	assert.NoError(err)

	want := []string{"url", "headers", "args", "method", "origin", "body", "ip", "user-agent", "cookies"}
	resp := MakeResponse(req, want...)

	assert.Equal(req.URL.Path, resp.URL)
	assert.Equal(fmt.Sprintf("%v", Flatten(req.Header)), fmt.Sprintf("%v", resp.Headers))
	assert.Equal(fmt.Sprintf("%v", Flatten(req.URL.Query())), fmt.Sprintf("%v", resp.Args))
	assert.Equal(req.Method, resp.Method)
	assert.Equal(req.Header.Get("origin"), resp.Origin)

	// assert body empty
	assert.Empty(resp.Data)
	assert.Empty(resp.Form)
	assert.Empty(resp.JSON)

	assert.Equal(strings.Split(req.RemoteAddr, ":")[0], resp.IP)
	assert.Equal(req.Header.Get("user-agent"), resp.UserAgent)

	for _, c := range req.Cookies() {
		val, ok := resp.Cookies[c.Name]
		assert.True(ok)
		assert.Equal(c.Value, val)
	}

}
