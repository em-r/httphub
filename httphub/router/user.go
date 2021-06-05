package router

import (
	"encoding/json"
	"net/http"

	"github.com/ElMehdi19/httphub/httphub/helpers"
)

func ViewUser(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /user User
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: The user's information.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - User
	//
	// responses:
	//   '200':
	//     description: The request's IP address, user agent and headers
	//
	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "ip", "user-agent", "headers"))
}

func ViewIP(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /ip User
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: The user's information.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - User
	//
	// responses:
	//   '200':
	//     description: The request's IP address
	//

	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "ip"))
}

func ViewUserAgent(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /user-agent User
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: The user's information.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - User
	//
	// responses:
	//   '200':
	//     description: The request's user-agent
	//

	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "user-agent"))
}
