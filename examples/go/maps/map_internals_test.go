package gomaps

import "fmt"

func addItem(m map[string]int) { m["x"] = 1 } // no pointer needed

func Example_internals() {
	m := make(map[string]int, 100) // size hint avoids regrowth
	addItem(m)
	fmt.Println(m)
	// Output: map[x:1]
}
