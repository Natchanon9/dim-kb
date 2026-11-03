package gocontext

import (
	"context"
	"fmt"
	"time"
)

func Example_withDeadline() {
	deadline := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	d, ok := ctx.Deadline()
	fmt.Println(d.Equal(deadline), ok)

	child, cancelChild := context.WithTimeout(ctx, time.Hour) // can't extend the parent
	defer cancelChild()
	cd, _ := child.Deadline()
	fmt.Println(cd.Equal(deadline))
	// Output:
	// true true
	// true
}
