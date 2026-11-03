//go:build go1.23

package moderngo

import (
	"fmt"
	"maps"
	"slices"
)

func Example_modernMaps() {
	m := map[string]int{"b": 2, "a": 1}
	fmt.Println(slices.Sorted(maps.Keys(m))) // the one-liner for sorted keys
	copyOfM := maps.Collect(maps.All(m))
	fmt.Println(len(copyOfM))
	// Output:
	// [a b]
	// 2
}
