package gocontext

import (
	"context"
	"errors"
	"fmt"
)

func Example_withCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()     // usually called by whoever decides to stop
	<-ctx.Done() // closed after cancel
	fmt.Println(ctx.Err(), errors.Is(ctx.Err(), context.Canceled))

	ctx2, cancel2 := context.WithCancelCause(context.Background()) // Go 1.20+
	cancel2(errors.New("user logged out"))
	fmt.Println(context.Cause(ctx2))
	// Output:
	// context canceled true
	// user logged out
}
