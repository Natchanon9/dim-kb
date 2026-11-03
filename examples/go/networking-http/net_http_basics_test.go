package networkinghttp

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func Example_basics() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
	})
	srv := httptest.NewServer(mux) // http.ListenAndServe(":8080", mux) in a real program
	defer srv.Close()

	res, err := http.Get(srv.URL + "/hello")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(res.StatusCode, string(body))
	// Output: 200 hello
}
