package goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// gen is a generator: it owns the channel and closes it when done.
func gen(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := range n {
			ch <- i
		}
	}()
	return ch
}

func Example_patterns() {
	var got []int
	for v := range gen(3) {
		got = append(got, v)
	}
	fmt.Println(got)

	sem := make(chan struct{}, 2) // at most 2 goroutines work at once
	var wg sync.WaitGroup
	var done atomic.Int32
	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}        // acquire
			defer func() { <-sem }() // release
			done.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println(done.Load())
	// Output:
	// [0 1 2]
	// 5
}
