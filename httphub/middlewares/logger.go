package middlewares

import (
	"log"
	"net/http"
	"strings"
)

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		path := r.URL.Path
		remote := strings.Split(r.RemoteAddr, ":")[0]
		log.Println(method, path, remote)

		h.ServeHTTP(w, r)
	})
}
