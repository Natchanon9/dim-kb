package standardlibrary

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

type person struct {
	Name string
	Age  int
}

func Example_sort() {
	ints := []int{3, 1, 2}
	sort.Ints(ints)
	fmt.Println(ints, sort.SearchInts(ints, 2))

	people := []person{{"b", 30}, {"a", 30}, {"c", 20}}
	sort.SliceStable(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println(people)

	slices.SortStableFunc(people, func(x, y person) int { return cmp.Compare(x.Name, y.Name) }) // modern style
	fmt.Println(people)
	// Output:
	// [1 2 3] 1
	// [{c 20} {b 30} {a 30}]
	// [{a 30} {b 30} {c 20}]
}
