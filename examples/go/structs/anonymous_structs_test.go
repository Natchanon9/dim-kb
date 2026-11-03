package structs

import "fmt"

func Example_anonymous() {
	point := struct{ X, Y int }{1, 2}
	tests := []struct {
		in   string
		want int
	}{
		{"a", 1},
		{"bb", 2},
	}
	fmt.Println(point, len(tests), tests[1].want)
	// Output: {1 2} 2 2
}
