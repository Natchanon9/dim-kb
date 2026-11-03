package moderngo

import (
	"cmp"
	"fmt"
	"slices"
)

type employee struct {
	Name string
	Age  int
}

func Example_cmp() {
	fmt.Println(cmp.Compare(1, 2), cmp.Less("b", "a"), cmp.Or("", "default")) // cmp.Or: first non-zero (1.22+)

	staff := []employee{{"bob", 30}, {"al", 30}, {"cy", 25}}
	slices.SortFunc(staff, func(a, b employee) int {
		return cmp.Or( // sort by age, then by name
			cmp.Compare(a.Age, b.Age),
			cmp.Compare(a.Name, b.Name),
		)
	})
	fmt.Println(staff)
	// Output:
	// -1 false default
	// [{cy 25} {al 30} {bob 30}]
}
