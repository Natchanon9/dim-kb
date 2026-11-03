package channels

import "fmt"

func numbers(done <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; ; i++ {
			select {
			case out <- i:
			case <-done: // caller lost interest
				return
			}
		}
	}()
	return out
}

func Example_done() {
	done := make(chan struct{})
	nums := numbers(done)
	fmt.Println(<-nums, <-nums)
	close(done) // broadcast: every receiver of done unblocks
	// Output: 0 1
}
