package pointers

import "fmt"

// Patch uses pointers so "not sent" (nil) differs from "set to zero".
type Patch struct {
	Name *string
	Age  *int
}

func Example_whenToUse() {
	age := 0
	patch := Patch{Age: &age}
	fmt.Println(patch.Name == nil, *patch.Age)
	// Output: true 0
}
