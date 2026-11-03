package concurrencypatterns

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Example_atomic() {
	var hits atomic.Int64
	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			hits.Add(1)
		}()
	}
	wg.Wait()

	var state atomic.Int32
	first := state.CompareAndSwap(0, 1)  // 0 → 1 succeeds
	second := state.CompareAndSwap(0, 1) // already 1: fails
	fmt.Println(hits.Load(), first, second)
	// Output: 1000 true false
}
