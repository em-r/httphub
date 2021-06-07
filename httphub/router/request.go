package router

import (
	"encoding/json"
	"net/http"

	"github.com/ElMehdi19/httphub/httphub/helpers"
)

func ViewRequest(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /request Request
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: The request's information.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Request inspection
	//
	// responses:
	//   '200':
	//     description: The request's IP address, user agent and headers
	//
	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "ip", "user-agent", "headers"))
}

func ViewIP(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /ip Request
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: The request's origin.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Request inspection
	//
	// responses:
	//   '200':
	//     description: The request's IP address
	//

	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "ip"))
}

func ViewUserAgent(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /user-agent Request
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: The request's user-agent.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Request inspection
	//
	// responses:
	//   '200':
	//     description: The request's user-agent
	//

	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "user-agent"))
}

func ViewHeaders(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /headers Request
	//
	// ---
	// produces:
	// - application/json
	//
	// summary: The request's headers.
	//
	// schemes:
	// - http
	// - https
	//
	// tags:
	// - Request inspection
	//
	// responses:
	//   '200':
	//     description: The request's headers
	//

	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "headers"))
}
