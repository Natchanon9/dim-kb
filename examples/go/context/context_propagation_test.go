package gocontext

import (
	"context"
	"fmt"
)

func repoFind(ctx context.Context) error {
	if err := ctx.Err(); err != nil { // check before expensive work
		return fmt.Errorf("repo: %w", err)
	}
	return nil
}

func serviceGet(ctx context.Context) error {
	if err := repoFind(ctx); err != nil { // pass the same ctx down
		return fmt.Errorf("service: %w", err)
	}
	return nil
}

func Example_propagation() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // e.g. the HTTP client disconnected
	fmt.Println(serviceGet(ctx))
	// Output: service: repo: context canceled
}
