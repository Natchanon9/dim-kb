package methodsinterfaces

import "fmt"

type Rect struct{ W, H float64 }

func (r Rect) Area() float64 { return r.W * r.H }

type Celsius float64 // methods work on any named type, not only structs

func (c Celsius) Fahrenheit() float64 { return float64(c)*9/5 + 32 }

func Example_methods() {
	fmt.Println(Rect{3, 4}.Area(), Celsius(100).Fahrenheit())
	// Output: 12 212
}
