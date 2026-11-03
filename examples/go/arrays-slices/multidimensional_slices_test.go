package arraysslices

import "fmt"

func Example_multidimensional() {
	grid := make([][]int, 2)
	for i := range grid {
		grid[i] = make([]int, 3) // each row is its own slice
	}
	grid[1][2] = 9
	fmt.Println(grid)
	// Output: [[0 0 0] [0 0 9]]
}
