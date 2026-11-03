package goroutines

import (
	"fmt"
	"sync"
)

func Example_basics() {
	var wg sync.WaitGroup
	results := make([]int, 3)
	for i := range 3 {
		wg.Add(1)
		go func() { // each goroutine writes its own index: no data race
			defer wg.Done()
			results[i] = i * i
		}()
	}
	wg.Wait()
	fmt.Println(results)
	// Output: [0 1 4]
}
