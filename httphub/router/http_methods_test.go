package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/ElMehdi19/httphub/httphub/helpers"
	"github.com/ElMehdi19/httphub/httphub/structs"
	"github.com/stretchr/testify/assert"
)

func TestMethodGetHandler(t *testing.T) {
	assert := assert.New(t)
	base := "http://127.0.0.1:5000/get"
	tcs := []structs.HTTPMethodsTestCase{
		{
			Name:    "basic",
			Args:    map[string][]string{},
			Headers: map[string][]string{},
		},
		{
			Name:    "with args",
			Args:    map[string][]string{"x": {"1", "2"}, "y": {"3"}},
			Headers: map[string][]string{},
		},
		{
			Name: "with headers",
			Args: map[string][]string{},
			Headers: map[string][]string{
				"scranton": {"bears", "beats", "battlestar galactica"},
				"whomai":   {"mehdi"},
				"origin":   {"https://mehdi.codes"},
			},
		},
		{
			Name: "default",
			Args: map[string][]string{"x": {"1", "2"}, "y": {"3"}},
			Headers: map[string][]string{
				"scranton": {"bears", "beats", "battlestar galactica"},
				"whomai":   {"mehdi"},
				"origin":   {"https://mehdi.codes"},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
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
				assert.True(reflect.DeepEqual(body.Args, tc.Args))
			}

			assert.True(reflect.DeepEqual(body.Headers, tc.Headers))
			assert.Empty(body.JSON)
		})
	}
}
