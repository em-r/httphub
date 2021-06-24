package router

import (
	"net/http"
	"net/url"

	"github.com/ElMehdi19/httphub/httphub/helpers"
	"github.com/gorilla/mux"
)

func viewRedirect(w http.ResponseWriter, r *http.Request, to string) {
	http.Redirect(w, r, to, http.StatusFound)
}

func ViewRedirect(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /redirect/{url} Redirects
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

func ViewRedirectRandom(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /redirect Redirects
	//
	// ---
	// summary: Redirects to a random relatice url.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Redirects
	//
	// responses:
	//   '302':
	//     description: Redirection to a random URL.

	to := helpers.Choose(topLevelGetpaths)
	viewRedirect(w, r, to)
}
