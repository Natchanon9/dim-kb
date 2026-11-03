package gocontext

import (
	"context"
	"fmt"
	"time"
)

func Example_basics() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	fmt.Println(slowOp(ctx)) // gives up after 10ms instead of 1s
	// Output: context deadline exceeded
}
