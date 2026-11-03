package bestpractices

import "fmt"

func Example_buildTags() {
	fmt.Println(Edition()) // `go test -tags premium` prints premium
	// Output: free
}
