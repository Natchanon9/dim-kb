package concurrencypatterns

import (
	"fmt"
	"slices"
	"sync"
)

func squareAll(jobs []int, workers int) []int {
	in := make(chan int)
	out := make(chan int, len(jobs))
	var wg sync.WaitGroup
	for range workers { // fixed number of goroutines
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range in {
				out <- j * j
			}
		}()
	}
	for _, j := range jobs {
		in <- j
	}
	close(in) // workers exit their range loop
	wg.Wait()
	close(out)

	var res []int
	for r := range out {
		res = append(res, r)
	}
	slices.Sort(res)
	return res
}

func Example_workerPool() {
	fmt.Println(squareAll([]int{1, 2, 3, 4}, 2))
	// Output: [1 4 9 16]
}
