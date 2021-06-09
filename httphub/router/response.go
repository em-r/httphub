package router

import (
	"encoding/json"
	"net/http"
)

func ViewResponseHeader(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /response-headers Response
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: Returns the passed query string args as headers.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Response inspection
	//
	// responses:
	//   '200':
	//     description: Response headers.

	args := r.URL.Query()
	for key, val := range args {
		if len(val) == 1 {
			w.Header().Add(key, val[0])
			continue
		}

		for _, nestedVal := range val {
			w.Header().Add(key, nestedVal)
		}
	}

	json.NewEncoder(w).Encode(w.Header())
}
