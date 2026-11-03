package structs

import (
	"fmt"
	"reflect"
)

type Point struct{ X, Y int }

type Tagged struct {
	Name string
	Tags []string // a slice field makes the struct non-comparable
}

func Example_comparison() {
	fmt.Println(Point{1, 2} == Point{1, 2})
	seen := map[Point]bool{{1, 2}: true} // comparable structs can be map keys
	fmt.Println(seen[Point{1, 2}])

	a := Tagged{"x", []string{"go"}}
	b := Tagged{"x", []string{"go"}}
	fmt.Println(reflect.DeepEqual(a, b)) // a == b would not compile
	// Output:
	// true
	// true
	// true
}
