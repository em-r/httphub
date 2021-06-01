package middlewares

import "net/http"

// JSONContent sets the content-type header in the response object
// to application/json.
func JSONContent(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		h.ServeHTTP(w, r)
	})
}

// CORS sets cors-related headers in the response object.
func CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	})
}
