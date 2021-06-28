package middlewares

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setUp(middlewares ...mux.MiddlewareFunc) (string, func()) {
	mux := mux.NewRouter()
	mux.Use(middlewares...)
	fn := func(w http.ResponseWriter, r *http.Request) {}

	mux.HandleFunc("/", fn)
	mux.HandleFunc("/file.js", fn)
	mux.HandleFunc("/file.txt", fn)
	mux.HandleFunc("/file.yaml", fn)
	mux.HandleFunc("/swagger/", fn)
	mux.HandleFunc("/whatever", fn)
	mux.HandleFunc("/explicit", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/xml")
	})

	server := httptest.NewServer(mux)
	return server.URL, func() {
		server.Close()
	}
}

func TestCORS(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUp(CORS)
	defer tearDown()

	resp, err := http.Get(base)
	assert.NoError(err)
	assert.Equal("*", resp.Header.Get("Access-Control-Allow-Origin"))
}

func TestContentType(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUp(ContentType)
	defer tearDown()

	type testCase struct {
		endpoint, contentType string
	}

	tcs := []testCase{
		{endpoint: "/", contentType: "text/html"},
		{endpoint: "/swagger/", contentType: "text/html"},
		{endpoint: "/file.js", contentType: "text/javascript"},
		{endpoint: "/file.txt", contentType: "text/plain"},
		{endpoint: "/file.yaml", contentType: "application/yaml"},
		{endpoint: "/whatever", contentType: "application/json"},
		{endpoint: "/explicit", contentType: "application/xml"},
	}

	for _, tc := range tcs {
		t.Run(tc.endpoint, func(t *testing.T) {
			resp, err := http.Get(fmt.Sprintf("%s/%s", base, tc.endpoint))
			assert.NoError(err)
			assert.Equal(tc.contentType, resp.Header.Get("content-type"))
		})
	}
}
