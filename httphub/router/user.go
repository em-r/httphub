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
	//   '200'
	//
	e := json.NewEncoder(w)
	e.Encode(helpers.MakeResponse(r, "ip", "user-agent"))
}
