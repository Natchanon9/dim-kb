package gomaps

import (
	"fmt"
	"maps"
)

func Example_mapsPackage() {
	a := map[string]int{"x": 1}
	b := maps.Clone(a)
	b["y"] = 2
	fmt.Println(maps.Equal(a, b), len(a), len(b))

	maps.Copy(a, b) // copy b's entries into a
	maps.DeleteFunc(a, func(k string, v int) bool { return v > 1 })
	fmt.Println(a)
	// Output:
	// false 1 2
	// map[x:1]
}
