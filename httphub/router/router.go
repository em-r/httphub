package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	handler := mux.NewRouter()

	handler.HandleFunc("/get", MethodGet).Methods("GET")

	return handler
}
