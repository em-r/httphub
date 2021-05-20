package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ElMehdi19/httphub/httphub/structs"
)

func parseJSON(r io.Reader) (interface{}, error) {
	var jsonBody interface{}
	err := json.NewDecoder(r).Decode(&jsonBody)
	return jsonBody, err
}

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

	switch r.Header.Get("content-type") {
	case "application/json":
		if body, err := parseJSON(r.Body); err == nil {
			resp.JSON = body
		} else {
			resp.JSON = err.Error()
		}
	case "application/x-www-form-urlencoded":
		if err := r.ParseForm(); err == nil {
			resp.Form = r.PostForm
		} else {
			resp.Args = map[string][]string{"error": {err.Error()}}
		}
	default:
		body := bytes.NewBuffer(nil)
		io.Copy(body, r.Body)
		resp.Data = body.String()
	}

	return resp
}
