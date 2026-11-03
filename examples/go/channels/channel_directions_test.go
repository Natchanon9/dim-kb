package channels

import "fmt"

func produce(out chan<- int) { // send-only
	for i := range 3 {
		out <- i
	}
	close(out)
}

func consume(in <-chan int) int { // receive-only
	sum := 0
	for v := range in {
		sum += v
	}
	return sum
}

func Example_directions() {
	ch := make(chan int)
	go produce(ch)
	fmt.Println(consume(ch))
	// Output: 3
}
