package gomaps

import (
	"fmt"
	"slices"
)

func Example_iterating() {
	m := map[string]int{"b": 2, "a": 1, "c": 3}
	keys := make([]string, 0, len(m))
	for k := range m { // random order
		keys = append(keys, k)
	}
	slices.Sort(keys) // sort for a stable order
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
	// Output:
	// a 1
	// b 2
	// c 3
}
