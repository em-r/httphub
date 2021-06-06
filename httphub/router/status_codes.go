package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func viewStatusCodes(w http.ResponseWriter, r *http.Request, code int) {
	w.WriteHeader(code)
	redirect := "https://mehdi.codes"

	switch code {
	case http.StatusMovedPermanently, http.StatusFound, http.StatusSeeOther, http.StatusTemporaryRedirect:
		http.Redirect(w, r, redirect, code)
	case http.StatusUnauthorized:
		w.Header().Set("WWW-Authenticate", "Basic realm='Access to staging site'")
	case http.StatusPaymentRequired:
		fmt.Fprint(w, "You owe me one!")
	case http.StatusUnsupportedMediaType:
		fmt.Fprint(w, `{"message": "Client did not request a supported media type."}`)
	case http.StatusProxyAuthRequired:
		w.Header().Set("Proxy-Authenticate", "Basic realm='Access to internal site'")
	}
}

func ViewStatusCodes(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /status/{code} Status codes
	//
	// ---
	// produces:
	// - application/json
	// - text/plain
	//
	// summary: Return the given status code
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Status codes
	//
	// parameters:
	// - in: path
	//   name: code
	//   description: Status code
	//   required: true
	//
	// responses:
	//   '100':
	//     description: Information responses
	//   '200':
	//     description: Success
	//   '300':
	//     description: Redirection
	//   '400':
	//     description: Client errors
	//   '500':
	//     description: Server errors

	vars := mux.Vars(r)
	code, err := strconv.Atoi(vars["code"])
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid status code %s", vars["code"]), http.StatusBadRequest)
		return
	}

	viewStatusCodes(w, r, code)
}
