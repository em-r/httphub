package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/ElMehdi19/httphub/httphub/helpers"
	"github.com/ElMehdi19/httphub/httphub/structs"
	"github.com/stretchr/testify/assert"
)

func TestMethodGetHandler(t *testing.T) {
	assert := assert.New(t)
	base := "http://127.0.0.1:5000/get"
	tc := structs.HTTPMethodsTestCase{
		Name: "default",
		Args: map[string][]string{"x": {"1", "2"}, "y": {"3"}},
		Headers: map[string][]string{
			"scranton": {"bears", "beats", "battlestar galactica"},
			"whomai":   {"mehdi"},
		},
	}

	url := helpers.CreateURL(base, tc.Args)
	req, err := http.NewRequest("GET", url, nil)
	req.Header = tc.Headers

	if !assert.NoError(err) {
		assert.FailNowf("could not create request: %s", err.Error())
	}
	rec := httptest.NewRecorder()
	MethodGet(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	assert.Equal("application/json", res.Header.Get("content-type"))

	var body structs.HTTPMethodsResponse
	if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
		assert.FailNowf("could not parse response body: %s", err.Error())
	}

	if len(tc.Args) == 0 {
		assert.Empty(body.Args)
	} else {
		assert.True(reflect.DeepEqual(tc.Args, body.Args))
	}

	assert.True(reflect.DeepEqual(tc.Headers, body.Headers))
	assert.Empty(body.JSON)
}

func TestMethodPostHandler(t *testing.T) {
	assert := assert.New(t)
	base := "http://127.0.0.1:5000/post"
	tcs := []structs.HTTPMethodsTestCase{
		{
			Name: "with json",
			Args: map[string][]string{"x": {"1", "2"}, "y": {"3"}},
			Headers: map[string][]string{
				"scranton": {"bears", "beats", "battlestar galactica"},
				"whomai":   {"mehdi"},
			},
			JSON: map[string]interface{}{
				"bool": true,
				"int":  1,
				"str":  "whatever",
			},
			ContentType: "application/json",
		},
		{
			Name: "with form",
			Args: map[string][]string{"x": {"1", "2"}, "y": {"3"}},
			Headers: map[string][]string{
				"scranton": {"bears", "beats", "battlestar galactica"},
				"whomai":   {"mehdi"},
			},
			Form: map[string][]string{
				"bool": {"true"},
				"int":  {"1"},
				"str":  {"whatever"},
			},
			ContentType: "application/x-www-form-urlencoded",
		},
	}

	for _, tc := range tcs {
		u := helpers.CreateURL(base, tc.Args)

		var b []byte
		switch tc.ContentType {
		case "application/json":
			b, _ = json.Marshal(tc.JSON)
		case "application/x-www-form-urlencoded":
			b = []byte(url.Values(tc.Form).Encode())
		default:
			b = nil
		}
		req, err := http.NewRequest("POST", u, bytes.NewReader(b))
		req.Header = tc.Headers
		req.Header.Add("content-type", tc.ContentType)

		if !assert.NoError(err) {
			assert.FailNow(err.Error())
		}

		rec := httptest.NewRecorder()
		MethodPost(rec, req)

		res := rec.Result()

		var body structs.HTTPMethodsResponse
		if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
			assert.FailNow(err.Error())
		}

		switch tc.ContentType {
		case "application/json":
			assert.Equal(fmt.Sprintf("%v", tc.JSON), fmt.Sprintf("%v", body.JSON))
		case "application/x-www-form-urlencoded":
			assert.Equal(fmt.Sprintf("%v", tc.Form), fmt.Sprintf("%v", body.Form))
		}

		if len(tc.Args) == 0 {
			assert.Empty(body.Args)
		} else {
			assert.True(reflect.DeepEqual(tc.Args, body.Args))
		}

		assert.True(reflect.DeepEqual(tc.Headers, body.Headers), body.Headers)
	}
}
