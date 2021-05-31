package middlewares

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setUp(middlewares ...mux.MiddlewareFunc) (string, func()) {
	mux := mux.NewRouter()
	mux.Use(middlewares...)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
			"success": true
		}`))
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

func TestJSONContent(t *testing.T) {
	assert := assert.New(t)
	base, tearDown := setUp(JSONContent)
	defer tearDown()

	resp, err := http.Get(base)
	assert.NoError(err)
	assert.Equal("application/json", resp.Header.Get("Content-Type"))

	var body map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		s, ok := body["success"]
		assert.True(ok)
		assert.True(s.(bool))
	}
}
