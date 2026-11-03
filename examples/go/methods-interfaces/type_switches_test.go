package methodsinterfaces

import "fmt"

func area(s any) float64 {
	switch v := s.(type) {
	case Rect:
		return v.Area()
	case *Square:
		return v.Area()
	default:
		return 0
	}
}

func Example_typeSwitches() {
	fmt.Println(area(Rect{2, 3}), area(&Square{3}), area("x"))
	// Output: 6 9 0
}
