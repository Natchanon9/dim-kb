package functions

import (
	"fmt"
	"sort"
)

func Example_anonymous() {
	result := func(a, b int) int { return a * b }(3, 4) // defined and called at once

	names := []string{"bob", "al", "christine"}
	sort.Slice(names, func(i, j int) bool { return len(names[i]) < len(names[j]) })
	fmt.Println(result, names)
	// Output: 12 [al bob christine]
}
