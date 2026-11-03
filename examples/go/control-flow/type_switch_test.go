package controlflow

import "fmt"

func describe(v any) string {
	switch x := v.(type) {
	case int:
		return fmt.Sprintf("int %d", x+1) // x is an int here
	case string:
		return "string " + x
	case nil:
		return "nil"
	default:
		return fmt.Sprintf("other %T", x)
	}
}

func Example_typeSwitch() {
	fmt.Println(describe(1))
	fmt.Println(describe("go"))
	fmt.Println(describe(nil))
	fmt.Println(describe(2.5))
	// Output:
	// int 2
	// string go
	// nil
	// other float64
}
