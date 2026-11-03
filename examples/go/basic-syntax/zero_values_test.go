package basicsyntax

import "fmt"

func Example_zeroValues() {
	var i int
	var s string
	var ok bool
	var p *int
	var list []int
	var m map[string]int
	fmt.Printf("%d %q %t %v %t %t\n", i, s, ok, p, list == nil, m == nil)
	// Output: 0 "" false <nil> true true
}
