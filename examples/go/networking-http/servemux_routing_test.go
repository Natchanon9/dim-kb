package networkinghttp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func Example_routing() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "user ", r.PathValue("id"))
	})
	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "created")
	})
	mux.HandleFunc("GET /files/{path...}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "file ", r.PathValue("path"))
	})

	do := func(method, path string) {
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, httptest.NewRequest(method, path, nil))
		fmt.Println(rec.Code, rec.Body.String())
	}
	do("GET", "/users/42")
	do("POST", "/users")
	do("GET", "/files/a/b.txt")
	do("DELETE", "/users/42")
	// Output:
	// 200 user 42
	// 201 created
	// 200 file a/b.txt
	// 405 Method Not Allowed
}
