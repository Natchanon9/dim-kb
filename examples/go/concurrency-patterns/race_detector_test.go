package concurrencypatterns

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Compiled but never run: `go test -race` would report this data race.
func Example_racy() {
	count := 0
	var wg sync.WaitGroup
	for range 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++ // unsynchronised read-modify-write
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func Example_raceFixed() {
	var count atomic.Int64
	var wg sync.WaitGroup
	for range 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println(count.Load())
	// Output: 2
}
