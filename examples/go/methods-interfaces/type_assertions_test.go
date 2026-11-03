package methodsinterfaces

import "fmt"

func Example_typeAssertions() {
	var v any = "go"
	s := v.(string)  // panics if v is not a string
	n, ok := v.(int) // safe form: ok reports success
	fmt.Println(s, n, ok)
	// Output: go 0 false
}
