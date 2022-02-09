package cors

import (
	"net/http"
)

func CorsMiddlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do stuff before intended handler here
		w.Header().Add("Access-Control-Allow-Origin", "*") // allow any origin
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-methods", "POST, HET, OPTIONS, PUT, DELETE") // allow http methods
		w.Header().Set("Access-Control-Allow-Headers", "*")                               // allow what headers

		handler.ServeHTTP(w, r)
		// do stuff after intended handler here
	})
}
