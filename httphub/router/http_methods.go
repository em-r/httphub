package router

import (
	"encoding/json"
	"net/http"

	"github.com/ElMehdi19/httphub/httphub/helpers"
)

// ViewBase is a view function that handles 'http_methods' routes
// e.g. /get /post ...
func ViewBase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r))
}
