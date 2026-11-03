package networkinghttp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

type greeter struct{ greeting string } // dependencies live in the struct

func (g greeter) ServeHTTP(w http.ResponseWriter, r *http.Request) { // implements http.Handler
	fmt.Fprintf(w, "%s, %s", g.greeting, r.URL.Query().Get("name"))
}

func Example_handlers() {
	rec := httptest.NewRecorder()
	greeter{"hi"}.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/?name=ann", nil))
	fmt.Println(rec.Body.String())
	// Output: hi, ann
}
