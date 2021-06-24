package router

import (
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func viewRedirect(w http.ResponseWriter, r *http.Request, to string) {
	http.Redirect(w, r, to, http.StatusFound)
}

func ViewRedirect(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /redirect/{url} Auth
	//
	// ---
	// summary: Redirects to the provided url.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Redirects
	//
	// parameters:
	// - in: path
	//   name: url
	//   description: username
	//   required: false
	//
	// responses:
	//   '302':
	//     description: Redirection to the specified URL.

	v := mux.Vars(r)
	to := url.URL{
		Scheme: "http",
		Host:   v["to"],
	}
	viewRedirect(w, r, to.String())
}
