package controlflow

import "fmt"

func Example_breakContinue() {
	for i := range 10 {
		if i%2 == 0 {
			continue // skip even numbers
		}
		if i > 7 {
			break // stop the loop
		}
		fmt.Print(i, " ")
	}
	// Output: 1 3 5 7
}
