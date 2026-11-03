package gomaps

import "fmt"

func Example_sets() {
	seen := map[string]struct{}{} // struct{} uses no memory
	for _, w := range []string{"go", "rust", "go"} {
		seen[w] = struct{}{}
	}
	_, ok := seen["go"]
	fmt.Println(len(seen), ok)
	// Output: 2 true
}
