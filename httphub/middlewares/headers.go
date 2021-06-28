package middlewares

import (
	"net/http"
	"strings"
)

// ContentType tries to detect the appropriate content-type of the response body and sets it to
// the Content-Type header if detected successfully, otherwise falls back to application/json
// since most endpoints in the app are serving back JSON responses.
func ContentType(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if the request is accessing a static file (e.g. Swagger UI bundles...)
		file := strings.Split(r.URL.Path, ".")
		if len(file) > 1 {
			extensions := map[string]string{
				"css":  "text/css",
				"html": "text/html",
				"js":   "text/javascript",
				"txt":  "text/plain",
				"yaml": "application/yaml",
				"json": "application/json",
			}
			ext := file[len(file)-1]
			if val, ok := extensions[ext]; ok {
				w.Header().Set("content-type", val)
				h.ServeHTTP(w, r)
				return
			}
		}

		var contentType string
		switch {
		// check if the request is accessing Swagger UI entry point.
		case r.URL.Path == "/", r.URL.Path == "/swagger/":
			contentType = "text/html"
		// falling back to application/json for the rest of the endpoints.
		case w.Header().Get("content-type") == "":
			contentType = "application/json"
		}

		w.Header().Set("content-type", contentType)
		h.ServeHTTP(w, r)
	})
}

// CORS sets cors-related headers in the response object.
func CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, PATCH, OPTIONS, HEAD")
		h.ServeHTTP(w, r)
	})
}
