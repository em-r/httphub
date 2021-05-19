package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/ElMehdi19/httphub/httphub/structs"
)

// MakeResponse creates and returns
// a structs.HTTPMethodsResponse instance
// based on the passed http.Request object.
func MakeResponse(r *http.Request) structs.HTTPMethodsResponse {
	resp := structs.HTTPMethodsResponse{
		URL:     r.URL.String(),
		Args:    r.URL.Query(),
		Headers: r.Header,
		Origin:  r.Header.Get("origin"),
	}
	if r.Method == "GET" {
		return resp
	}

	var jsonBody interface{}
	if err := json.NewDecoder(r.Body).Decode(&jsonBody); err == nil {
		resp.JSON = jsonBody
	} else {
		resp.JSON = err.Error()
	}
	return resp
}
