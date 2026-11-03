package concurrencypatterns

import (
	"fmt"
	"sync"
)

func Example_once() {
	var once sync.Once
	calls := 0
	for range 3 {
		once.Do(func() { calls++ }) // runs only the first time
	}

	config := sync.OnceValue(func() string { // Go 1.21+: once + cached result
		return "loaded"
	})
	fmt.Println(calls, config(), config())
	// Output: 1 loaded loaded
}
