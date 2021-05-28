package router

import (
	"encoding/json"
	"net/http"

	"github.com/ElMehdi19/httphub/httphub/helpers"
)

// ViewAny is a view function that handles all sorts of http requests.
func ViewAny(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "url", "headers", "args", "body", "method"))
}

// ViewAll is a view function that handles GET requests.
func ViewGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "url", "headers", "args"))
}

// ViewAll is a view function that handles POST requests.
func ViewPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "url", "headers", "args", "body"))
}

// ViewAll is a view function that handles PUT requests.
func ViewPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "url", "headers", "args", "body"))
}

// ViewAll is a view function that handles PATCH requests.
func ViewPatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "url", "headers", "args", "body"))
}

// ViewAll is a view function that handles DELETE requests.
func ViewDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "url", "headers", "args", "body"))
}
