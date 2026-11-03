package concurrencypatterns

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu sync.Mutex
	m  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key]++
}

func (c *SafeCounter) Get(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.m[key]
}

func Example_mutex() {
	c := SafeCounter{m: map[string]int{}}
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc("hits")
		}()
	}
	wg.Wait()
	fmt.Println(c.Get("hits"))
	// Output: 100
}
