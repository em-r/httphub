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

	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "cookies"))
}

func ViewSetCookies(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /cookies/set Cookies
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: Sets cookie data and redirects to /cookies.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Cookies
	//
	// responses:
	//   '302':
	//     description: Sets cookie data and redirects to /cookies.

	for key, vals := range r.URL.Query() {
		c := http.Cookie{
			Name:  key,
			Value: vals[0],
		}
		http.SetCookie(w, &c)
	}
	http.Redirect(w, r, "/cookies", http.StatusFound)
}