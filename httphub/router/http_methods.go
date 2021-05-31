package router

import (
	"encoding/json"
	"net/http"

	"github.com/ElMehdi19/httphub/httphub/helpers"
)

// ViewAny is a view function that handles all sorts of http requests.
func ViewAny(w http.ResponseWriter, r *http.Request) {
	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "url", "headers", "args", "body", "method"))
}

// ViewAll is a view function that handles GET requests.
func ViewGet(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /get getRequest
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: The request's query args.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - HTTP Methods
	//
	// responses:
	//   '200':
	//     description: The request's query args.

	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "url", "headers", "args"))
}

// ViewAll is a view function that handles POST requests.
func ViewPost(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /post postRequest
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: The request's POST params.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - HTTP Methods
	//
	// responses:
	//   '200':
	//     description: The request's post params.

	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "url", "headers", "args", "body"))
}

// ViewAll is a view function that handles PUT requests.
func ViewPut(w http.ResponseWriter, r *http.Request) {
	// swagger:operation PUT /put postRequest
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: The request's PUT params.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - HTTP Methods
	//
	// responses:
	//   '200':
	//     description: The request's PUT params.

	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "url", "headers", "args", "body"))
}

// ViewAll is a view function that handles PATCH requests.
func ViewPatch(w http.ResponseWriter, r *http.Request) {
	// swagger:operation PATCH /patch postRequest
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: The request's PATCH params.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - HTTP Methods
	//
	// responses:
	//   '200':
	//     description: The request's PATCH params.

	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "url", "headers", "args", "body"))
}

// ViewAll is a view function that handles DELETE requests.
func ViewDelete(w http.ResponseWriter, r *http.Request) {
	// swagger:operation DELETE /delete postRequest
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: The request's DELETE params.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - HTTP Methods
	//
	// responses:
	//   '200':
	//     description: The request's DELETE params.

	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "url", "headers", "args", "body"))
}
