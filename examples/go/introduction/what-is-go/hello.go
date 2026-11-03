// Package whatisgo shows why Go fits backend work: a working HTTP API using
// only the standard library, summarised in src/content/docs/programming/go/introduction/what-is-go.md.
package whatisgo

import (
	"encoding/json"
	"net/http"
)

// NewMux returns a router with one route; method + path params need Go 1.22+.
func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "hello " + r.PathValue("name"),
		})
	})
	return mux
}
