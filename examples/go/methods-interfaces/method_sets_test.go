package methodsinterfaces

import "fmt"

type Shape interface{ Area() float64 }

type Square struct{ S float64 }

func (s *Square) Area() float64 { return s.S * s.S } // pointer receiver

func Example_methodSets() {
	var sh Shape = &Square{2} // Square{2} would not compile: only *Square has Area
	fmt.Println(sh.Area())
	// Output: 4
}
