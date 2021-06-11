package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ElMehdi19/httphub/httphub/helpers"
	"github.com/gorilla/mux"
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

func ViewCache(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /cache Response
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: Returns a 304 if an If-Modified-Since header or If-None-Match is present. Otherwise returns 200.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Response inspection
	//
	// parameters:
	// - in: headers
	//   name: if-Modified-Since
	//   required: false
	//
	// - in: headers
	//   name: If-None-Match
	//   required: false
	//
	// responses:
	//   '200':
	//     description: Cached response.
	//   '304':
	//     description: Modified.

	if r.Header.Get("If-Modified-Since") == "" && r.Header.Get("If-None-Match") == "" {
		httpDate := time.Now().Format("Mon 02-Jan-2006 15:04:05 GMT")
		w.Header().Set("Date", httpDate)
		w.Header().Set("Content-Location", r.URL.Path)
		w.Header().Set("Etag", helpers.RandomStr(23))
		ViewGet(w, r)
		return
	}

	viewStatusCodes(w, r, http.StatusNotModified)
}

func viewCacheControl(w http.ResponseWriter, r *http.Request, maxAge string) {
	seconds, err := strconv.Atoi(maxAge)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("cache-control", "public")
	w.Header().Add("cache-control", fmt.Sprintf("max-age=%d", seconds))
	ViewGet(w, r)
}

func ViewCacheControl(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /cache/{value} Response
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: Sets Cache-Control header for n seconds.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Response inspection
	//
	// parameters:
	// - in: path
	//   name: value
	//   required: true
	//   type: integer
	//
	// responses:
	//   '200':
	//     description: Cached response.
	//   '400':
	//     description: Path variable value not an integer.

	v := mux.Vars(r)
	maxAge := v["value"]
	viewCacheControl(w, r, maxAge)
}

func ViewJSONResponse(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /json Response
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: Returns a JSON document.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Response formats
	//
	// responses:
	//   '200':
	//     description: JSON document.

	w.Write([]byte(helpers.JSONDoc))
}
