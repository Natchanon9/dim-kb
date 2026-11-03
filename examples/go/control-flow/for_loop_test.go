package controlflow

import "fmt"

func Example_forLoop() {
	sum := 0
	for i := 0; i < 5; i++ { // classic three-part loop
		sum += i
	}
	n := 1
	for n < 100 { // "while" loop
		n *= 2
	}
	count := 0
	for { // infinite loop
		count++
		if count == 3 {
			break
		}
	}
	for i := range 3 { // Go 1.22+: range over an int
		fmt.Print(i, " ")
	}
	fmt.Println(sum, n, count)
	// Output: 0 1 2 10 128 3
}
