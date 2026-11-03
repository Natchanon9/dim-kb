package methodsinterfaces

import "fmt"

func Example_emptyInterface() {
	var v any = 42 // any == interface{}
	v = "hi"
	v = []int{1}
	fmt.Printf("%T\n", v)

	values := []any{1, "a", true, nil}
	fmt.Println(values...)
	// Output:
	// []int
	// 1 a true <nil>
}
