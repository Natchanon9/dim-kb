package networkinghttp

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

type ridKey struct{}

func withRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ridKey{}, "req-7")
		next.ServeHTTP(w, r.WithContext(ctx)) // pass the enriched request down
	})
}

func report(w http.ResponseWriter, r *http.Request) {
	select {
	case <-time.After(time.Second): // slow work
		fmt.Fprint(w, "done")
	case <-r.Context().Done(): // client went away or server is shutting down
		fmt.Println(r.Context().Value(ridKey{}), "stopped:", r.Context().Err())
	}
}

func Example_requestContext() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // simulate a client that disconnected
	req := httptest.NewRequest(http.MethodGet, "/report", nil).WithContext(ctx)
	withRequestID(http.HandlerFunc(report)).ServeHTTP(httptest.NewRecorder(), req)
	// Output: req-7 stopped: context canceled
}
