package router

import (
	"encoding/json"
	"net/http"

	"github.com/ElMehdi19/httphub/httphub/helpers"
)

func ViewCookies(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /cookies Cookies
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: Cookie data.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Cookies
	//
	// responses:
	//   '200':
	//     description: Cookie data.
	//
	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "cookies"))
}
