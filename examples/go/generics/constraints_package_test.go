package generics

import (
	"cmp"
	"fmt"
	"slices"
)

func Index[T comparable](s []T, v T) int { // comparable: == and !=
	for i, x := range s {
		if x == v {
			return i
		}
	}
	return -1
}

func SortedKeys[K cmp.Ordered, V any](m map[K]V) []K { // cmp.Ordered: < > <= >=
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

func Example_constraintsPackage() {
	fmt.Println(Index([]string{"a", "b"}, "b"), SortedKeys(map[string]int{"b": 1, "a": 2}))
	// Output: 1 [a b]
}
