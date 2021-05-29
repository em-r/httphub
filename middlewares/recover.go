package middlewares

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/ElMehdi19/httphub/httphub/helpers"
)

func Recover(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if !helpers.IsDevMode() {
					http.Error(w, `{
						"error": true
					}`, http.StatusInternalServerError)
					return
				}

				stack := debug.Stack()
				fmt.Fprintf(w, "<h1>Panic: %v</h1><pre>%s</pre>", err, string(stack))
				return
			}
		}()
		h.ServeHTTP(w, r)
	})
}
