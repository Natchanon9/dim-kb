package networkinghttp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("->", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func requireKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-Key") != "secret" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return // stop the chain
		}
		next.ServeHTTP(w, r)
	})
}

func Example_middleware() {
	ok := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "data") })
	h := logging(requireKey(ok)) // outermost runs first

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/data", nil))
	fmt.Println(rec.Code)
	// Output:
	// -> GET /data
	// 401
}
