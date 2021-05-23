package router

import (
	"encoding/json"
	"net/http"

	"github.com/ElMehdi19/httphub/httphub/helpers"
)

// MethodGet handles the requests hitting
// the `/get` endpoint and sends back
// a JSONed structs.HTTPMethodsResponse.
func MethodGET(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r))
}

func MethodPOST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r))
}
