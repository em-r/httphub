package router

import (
	"encoding/json"
	"net/http"

	"github.com/ElMehdi19/httphub/httphub/structs"
	"github.com/gorilla/mux"
)

func viewBasicAuth(w http.ResponseWriter, r *http.Request, user, passwd string) {
	username, password, ok := r.BasicAuth()

	if !ok || user != username || passwd != password {
		w.Header().Set("WWW-Authenticate", "Basic realm=Sign in")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	resp := structs.AuthResponse{
		Authorized: true,
		User:       user,
	}

	json.NewEncoder(w).Encode(&resp)
}

func ViewBasicAuth(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /auth/basic/{user}/{passwd} Auth
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: Basic Auth protected route.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Auth
	//
	// parameters:
	// - in: path
	//   name: user
	//   description: username
	//   required: false
	//
	// - in: path
	//   name: passwd
	//   description: password
	//   required: false
	//
	// responses:
	//   '200':
	//     description: Successful authentication
	//   '401':
	//     description: Unsuccessful authentication

	v := mux.Vars(r)
	user := v["user"]
	passwd := v["passwd"]

	viewBasicAuth(w, r, user, passwd)
}
