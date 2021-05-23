package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	handler := mux.NewRouter()

	handler.HandleFunc("/get", MethodGET).Methods("GET")
	handler.HandleFunc("/post", MethodPOST).Methods("POST")

	return handler
}
