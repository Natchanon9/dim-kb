package controlflow

import "fmt"

func Example_range() {
	for i, v := range []string{"a", "b"} {
		fmt.Println(i, v)
	}
	for i, r := range "gö" { // i is the byte offset, r is a rune
		fmt.Println(i, string(r))
	}
	// Output:
	// 0 a
	// 1 b
	// 0 g
	// 1 ö
}
