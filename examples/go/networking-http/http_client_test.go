package networkinghttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func Example_client() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"id":1,"name":"ann"}`)
	}))
	defer srv.Close()

	client := &http.Client{Timeout: 5 * time.Second} // never use a client without a timeout
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL, nil)
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() // always close, even if you don't read it
	if res.StatusCode != http.StatusOK {
		panic(res.Status) // a 4xx/5xx is not an error from Do
	}
	var user struct {
		ID   int
		Name string
	}
	json.NewDecoder(res.Body).Decode(&user)
	fmt.Println(user.ID, user.Name)
	// Output: 1 ann
}
