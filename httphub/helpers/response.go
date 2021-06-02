package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/ElMehdi19/httphub/httphub/structs"
)

func parseJSON(r io.Reader) (interface{}, error) {
	var jsonBody interface{}
	err := json.NewDecoder(r).Decode(&jsonBody)
	return jsonBody, err
}

func parseBody(r *http.Request, resp *structs.HTTPMethodsResponse) {
	if r.Method == "GET" {
		return
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
			resp.Form = Flatten(r.PostForm)
		} else {
			log.Println(err.Error())
		}
	default:
		body := bytes.NewBuffer(nil)
		io.Copy(body, r.Body)
		// resp.Data = body.Sting()
		if body.String() != "" {
			resp.Data = body.String()
		}
	}
}

// MakeResponse creates and returns a structs.HTTPMethodsResponse
// instance populated with the field names passed on the want variadic param.
func MakeResponse(r *http.Request, want ...string) structs.HTTPMethodsResponse {
	var resp structs.HTTPMethodsResponse
	keys := []string{"url", "headers", "args", "method", "body", "origin", "form", "ip", "user-agent"}
	isValid := func(field string) bool {
		for _, key := range keys {
			if key == field {
				return true
			}
		}
		return false
	}

	for _, f := range want {
		if !isValid(f) {
			continue
		}

		switch f {
		case "url":
			resp.URL = r.URL.Path
		case "headers":
			resp.Headers = Flatten(r.Header)
		case "args":
			resp.Args = Flatten(r.URL.Query())
		case "method":
			resp.Method = r.Method
		case "origin":
			resp.Origin = r.Header.Get("origin")
		case "body":
			parseBody(r, &resp)
		case "ip":
			resp.IP = strings.Split(r.RemoteAddr, ":")[0]
		case "user-agent":
			resp.UserAgent = r.Header.Get("user-agent")
		}
	}

	return resp
}
