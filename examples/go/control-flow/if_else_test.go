package controlflow

import "fmt"

func grade(score int) string {
	if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B"
	}
	return "C"
}

func Example_ifElse() {
	fmt.Println(grade(85), grade(72), grade(10))
	// Output: A B C
}
