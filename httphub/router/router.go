package router

import (
	"net/http"

	"github.com/ElMehdi19/httphub/httphub/helpers"
	"github.com/ElMehdi19/httphub/middlewares"
	"github.com/gorilla/mux"
)

func New() http.Handler {
	handler := mux.NewRouter()
	handler.Use(middlewares.Recover, middlewares.Logger, middlewares.JSONContent)

	handler.HandleFunc("/get", ViewGet).Methods("GET")
	handler.HandleFunc("/put", ViewPut).Methods("PUT")
	handler.HandleFunc("/post", ViewPost).Methods("POST")
	handler.HandleFunc("/patch", ViewPatch).Methods("patch")
	handler.HandleFunc("/delete", ViewDelete).Methods("delete")
	handler.HandleFunc("/any", ViewAny)

	if helpers.IsDevMode() {
		handler.HandleFunc("/debug", SourceCodeHandler)
	}

	return handler
}
