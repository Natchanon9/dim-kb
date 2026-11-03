package concurrencypatterns

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func fetch(ctx context.Context, url string) error {
	if url == "bad" {
		return fmt.Errorf("fetch %s: failed", url)
	}
	select {
	case <-ctx.Done(): // another goroutine failed: stop early
		return ctx.Err()
	default:
		return nil
	}
}

func Example_errgroup() {
	g, ctx := errgroup.WithContext(context.Background())
	g.SetLimit(2) // at most 2 goroutines at once
	for _, u := range []string{"a", "bad", "c"} {
		g.Go(func() error { return fetch(ctx, u) })
	}
	fmt.Println(g.Wait()) // first error; ctx is cancelled when it happens
	// Output: fetch bad: failed
}
