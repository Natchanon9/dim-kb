package gomaps

import "fmt"

func Example_mapOperations() {
	m := map[string]int{"a": 1, "b": 2}
	m["a"] = 10       // update
	delete(m, "b")    // remove
	delete(m, "nope") // no-op, no error
	fmt.Println(m, len(m))
	clear(m) // Go 1.21+: remove everything
	fmt.Println(len(m))
	// Output:
	// map[a:10] 1
	// 0
}
