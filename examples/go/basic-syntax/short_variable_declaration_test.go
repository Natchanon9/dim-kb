package basicsyntax

import "fmt"

func Example_shortDeclaration() {
	count := 10
	name, ok := "go", true
	count, extra := 20, 1 // allowed: extra is new, count is reassigned
	fmt.Println(count, name, ok, extra)
	// Output: 20 go true 1
}
