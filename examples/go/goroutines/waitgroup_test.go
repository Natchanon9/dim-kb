package goroutines

import (
	"fmt"
	"sync"
)

func Example_waitGroup() {
	var (
		wg    sync.WaitGroup
		mu    sync.Mutex
		total int
	)
	for _, n := range []int{1, 2, 3} {
		wg.Add(1) // before starting the goroutine
		go func() {
			defer wg.Done()
			mu.Lock()
			total += n
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(total)
	// Output: 6
}
