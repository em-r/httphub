package router

import "net/http"

func MethodGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get"))
}
