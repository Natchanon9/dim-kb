package channels

import (
	"fmt"
	"sync"
)

func merge(chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	for _, c := range chs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}()
	}
	go func() { wg.Wait(); close(out) }() // close only after every input is drained
	return out
}

func Example_fanIn() {
	sum := 0
	for v := range merge(emit(1, 2), emit(3)) {
		sum += v
	}
	fmt.Println(sum)
	// Output: 6
}
