package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/ElMehdi19/httphub/httphub/helpers"
	"github.com/ElMehdi19/httphub/httphub/structs"
	"github.com/stretchr/testify/assert"
)

// testArgs tests if the request's query args
// are the same as those returned in the response
// body.
func testArgs(t *testing.T, tc structs.HTTPMethodsTestCase, body structs.HTTPMethodsResponse) {
	assert := assert.New(t)
	if len(tc.Args) == 0 {
		assert.Empty(body.Args)
	} else {
		assert.True(reflect.DeepEqual(tc.Args, body.Args))
	}
}

// testResponseBody tests if the request's body
// is the same as the appropriate body field in
// the response body. The field to be checked
// is determined based on the request's content-type
// header.
func testResponseBody(t *testing.T, tc structs.HTTPMethodsTestCase, body structs.HTTPMethodsResponse) {
	assert := assert.New(t)
	switch tc.ContentType {
	case "application/json":
		assert.Equal(fmt.Sprintf("%v", tc.JSON), fmt.Sprintf("%v", body.JSON))
	case "application/x-www-form-urlencoded":
		assert.Equal(fmt.Sprintf("%v", tc.Form), fmt.Sprintf("%v", body.Form))
	default:
		assert.Equal(tc.Data, body.Data)
	}
}

// testMethodWithBody is a proxy function used to test
// all the request's that hold a body (e.g. POST, PUT...).
// This is function is to be called inside the exported
// test functions that handle those kind of requests.
func testMethodWithBody(t *testing.T, tc structs.HTTPMethodsTestCase, method string) {
	assert := assert.New(t)
	base := fmt.Sprintf("http://127.0.0.1:5000/%s", method)

	u := helpers.CreateURL(base, tc.Args)
	b := helpers.MakeBodyFromTestCase(tc)

	req, err := http.NewRequest(method, u, bytes.NewReader(b))
	req.Header = tc.Headers
	req.Header.Set("content-type", tc.ContentType)

	if !assert.NoError(err) {
		assert.FailNow(err.Error())
	}

	rec := httptest.NewRecorder()
	MethodPOST(rec, req)

	res := rec.Result()

	var body structs.HTTPMethodsResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		assert.FailNow(err.Error())
	}

	testResponseBody(t, tc, body)
	testArgs(t, tc, body)

	assert.True(reflect.DeepEqual(tc.Headers, body.Headers), body.Headers)
	if body.Method != "" {
		assert.Equal(body.Method, req.Method)
	}
}

func TestMethodGetHandler(t *testing.T) {
	assert := assert.New(t)
	base := "http://127.0.0.1:5000/get"
	tc := helpers.HTTPMethodsBaseTc

	url := helpers.CreateURL(base, tc.Args)
	req, err := http.NewRequest("GET", url, nil)
	req.Header = tc.Headers

	if !assert.NoError(err) {
		assert.FailNowf("could not create request: %s", err.Error())
	}

	rec := httptest.NewRecorder()
	MethodGET(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	assert.Equal("application/json", res.Header.Get("content-type"))

	var body structs.HTTPMethodsResponse
	if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
		assert.FailNowf("could not parse response body: %s", err.Error())
	}

	// assert query args are the same as those in the response body.
	testArgs(t, tc, body)

	// assert request headers are the some as those returned in the response body.
	assert.True(reflect.DeepEqual(tc.Headers, body.Headers))
	// body fields in the response body must be empty.
	assert.Empty(body.JSON)
	assert.Empty(body.Form)
	assert.Empty(body.Data)
}

func TestMethodPostHandler(t *testing.T) {
	for _, tc := range helpers.HTTPMethodsTcs {
		t.Run(tc.Name, func(t *testing.T) {
			testMethodWithBody(t, tc, "POST")
		})
	}
}
