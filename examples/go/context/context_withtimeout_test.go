package gocontext

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func Example_withTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel() // always: releases the timer early
	err := slowOp(ctx)
	fmt.Println(errors.Is(err, context.DeadlineExceeded))
	// Output: true
}
