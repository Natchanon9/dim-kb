package channels

import "fmt"

func Example_range() {
	ch := make(chan string, 3)
	ch <- "a"
	ch <- "b"
	close(ch) // range stops after the buffered values are drained
	for s := range ch {
		fmt.Println(s)
	}
	v, ok := <-ch // closed and empty: zero value, ok=false
	fmt.Printf("%q %v\n", v, ok)
	// Output:
	// a
	// b
	// "" false
}
