package generics

import "fmt"

type Celsius float64

type Float interface {
	~float32 | ~float64 // ~ includes named types like Celsius
}

func Avg[T Float](xs ...T) T {
	var total T
	for _, x := range xs {
		total += x
	}
	return total / T(len(xs))
}

func Example_typeSets() {
	fmt.Println(Avg(Celsius(20), Celsius(30)))
	// Output: 25
}
