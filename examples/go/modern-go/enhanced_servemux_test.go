package moderngo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func Example_enhancedServeMux() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "home") }) // exactly "/"
	mux.HandleFunc("GET /posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "post ", r.PathValue("id"))
	})
	mux.HandleFunc("GET /posts/latest", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "latest") // more specific than /posts/{id}: wins
	})

	for _, p := range []string{"/", "/posts/7", "/posts/latest", "/nope"} {
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, httptest.NewRequest("GET", p, nil))
		fmt.Println(p, rec.Code, rec.Body.String())
	}
	// Output:
	// / 200 home
	// /posts/7 200 post 7
	// /posts/latest 200 latest
	// /nope 404 404 page not found
}
