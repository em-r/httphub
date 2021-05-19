package helpers

import (
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
	return resp
}
