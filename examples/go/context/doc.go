// Package gocontext holds the snippets for the "Context Package" pages
// in src/content/docs/programming/go/context/.
package gocontext

import (
	"context"
	"time"
)

// slowOp pretends to work for one second but stops as soon as ctx is done.
func slowOp(ctx context.Context) error {
	select {
	case <-time.After(time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
