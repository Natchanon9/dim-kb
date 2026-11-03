package goroutines

import (
	"context"
	"fmt"
)

func counter(ctx context.Context, out chan<- int) {
	for i := 0; ; i++ {
		select {
		case out <- i:
		case <-ctx.Done(): // without this case the goroutine blocks forever: a leak
			return
		}
	}
}

func Example_leaks() {
	ctx, cancel := context.WithCancel(context.Background())
	out := make(chan int)
	go counter(ctx, out)
	fmt.Println(<-out, <-out)
	cancel() // lets counter return
	// Output: 0 1
}
