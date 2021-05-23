package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	handler := mux.NewRouter()

	handler.HandleFunc("/get", ViewBase).Methods("GET")
	handler.HandleFunc("/put", ViewBase).Methods("PUT")
	handler.HandleFunc("/post", ViewBase).Methods("POST")
	handler.HandleFunc("/patch", ViewBase).Methods("patch")
	handler.HandleFunc("/delete", ViewBase).Methods("delete")

	return handler
}
