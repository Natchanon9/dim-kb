package goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Example_manyGoroutines() {
	var (
		wg sync.WaitGroup
		n  atomic.Int64
	)
	for range 10_000 { // 10k goroutines is cheap; 10k OS threads is not
		wg.Add(1)
		go func() {
			defer wg.Done()
			n.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println(n.Load())
	// Output: 10000
}
