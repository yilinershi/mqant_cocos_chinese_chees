//Package algoutil contain some scaffold algo
package algoutil

import (
	"encoding/json"
	"net/http"
)

func AccessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		h.ServeHTTP(w, r)
	})
}

func OptionControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			json.NewEncoder(w).Encode(`{ "code": 0, "data": "success"}`)
			return
		}

		h.ServeHTTP(w, r)
	})
}
