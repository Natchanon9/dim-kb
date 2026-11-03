package channels

import "fmt"

func Example_nilChannels() {
	a, b := emit(1), emit(2)
	sum := 0
	for a != nil || b != nil {
		select {
		case v, ok := <-a:
			if !ok {
				a = nil // a nil channel blocks forever: this case is now disabled
				continue
			}
			sum += v
		case v, ok := <-b:
			if !ok {
				b = nil
				continue
			}
			sum += v
		}
	}
	fmt.Println(sum)
	// Output: 3
}
