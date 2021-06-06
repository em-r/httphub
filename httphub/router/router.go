package router

import (
	"net/http"

	"github.com/ElMehdi19/httphub/httphub/helpers"
	"github.com/ElMehdi19/httphub/httphub/middlewares"
	"github.com/gorilla/mux"
)

func New() http.Handler {
	handler := mux.NewRouter()
	handler.Use(middlewares.Recover, middlewares.Logger, middlewares.JSONContent, middlewares.CORS)

	if helpers.IsDevMode() {
		handler.HandleFunc("/debug", SourceCodeHandler)
	}

	// HTTP Methods handlers
	handler.HandleFunc("/get", ViewGet).Methods("GET")
	handler.HandleFunc("/put", ViewPut).Methods("PUT")
	handler.HandleFunc("/post", ViewPost).Methods("POST")
	handler.HandleFunc("/patch", ViewPatch).Methods("patch")
	handler.HandleFunc("/delete", ViewDelete).Methods("delete")
	handler.HandleFunc("/any", ViewAny)

	// User Handlers
	handler.HandleFunc("/user", ViewUser).Methods("GET")
	handler.HandleFunc("/ip", ViewIP).Methods("GET")
	handler.HandleFunc("/user-agent", ViewUserAgent).Methods("GET")
	handler.HandleFunc("/headers", ViewHeaders).Methods("GET")

	// status codes handlers
	handler.HandleFunc("/status/{code}", ViewStatusCodes)

	return handler
}
