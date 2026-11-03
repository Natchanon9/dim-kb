package channels

import (
	"fmt"
	"slices"
	"sync"
)

func Example_fanOut() {
	jobs := emit(1, 2, 3, 4)
	results := make(chan int)
	var wg sync.WaitGroup
	for range 3 { // 3 workers read from one channel
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				results <- j * 10
			}
		}()
	}
	go func() { wg.Wait(); close(results) }()

	var got []int
	for r := range results {
		got = append(got, r)
	}
	slices.Sort(got) // arrival order is not deterministic
	fmt.Println(got)
	// Output: [10 20 30 40]
}
