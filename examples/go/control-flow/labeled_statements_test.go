package controlflow

import "fmt"

func Example_labels() {
outer:
	for i := range 3 {
		for j := range 3 {
			if j == 2 {
				continue outer // next i
			}
			if i == 2 {
				break outer // leave both loops
			}
			fmt.Printf("%d%d ", i, j)
		}
	}
	// Output: 00 01 10 11
}
