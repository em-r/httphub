package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViewStatusCodes(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/status", nil)
	assert.NoError(err)

	type testCase struct {
		code         int
		header, body string
	}

	tcs := []testCase{
		{
			code: http.StatusMovedPermanently,
		},
		{
			code:   http.StatusUnauthorized,
			header: "WWW-Authenticate",
		},
		{
			code: http.StatusPaymentRequired,
			body: "You owe me one!",
		},
		{
			code: http.StatusUnsupportedMediaType,
			body: `{"message": "Client did not request a supported media type."}`,
		},
		{
			code:   http.StatusProxyAuthRequired,
			header: "Proxy-Authenticate",
		},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%d", tc.code), func(t *testing.T) {
			handler := func(w http.ResponseWriter, r *http.Request) {
				viewStatusCodes(w, r, tc.code)
			}

			rec := httptest.NewRecorder()
			handler(rec, req)
			assert.Equal(tc.code, rec.Result().StatusCode)

			if tc.header != "" {
				assert.NotEmpty(rec.Header().Get(tc.header))
			}

			if tc.body != "" {
				assert.Equal(tc.body, rec.Body.String(), rec.Body.String())
			}
		})
	}
}
