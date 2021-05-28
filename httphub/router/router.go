package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	handler := mux.NewRouter()

	handler.HandleFunc("/get", ViewGet).Methods("GET")
	handler.HandleFunc("/put", ViewPut).Methods("PUT")
	handler.HandleFunc("/post", ViewPost).Methods("POST")
	handler.HandleFunc("/patch", ViewPatch).Methods("patch")
	handler.HandleFunc("/delete", ViewDelete).Methods("delete")
	handler.HandleFunc("/any", ViewAny)

	return handler
}
