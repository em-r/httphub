package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	handler := mux.NewRouter()

	handler.HandleFunc("/get", MethodGet).Methods("GET")
	handler.HandleFunc("/post", MethodPost).Methods("POST")

	return handler
}
