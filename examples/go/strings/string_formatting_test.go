package gostrings

import "fmt"

func Example_formatting() {
	fmt.Println(fmt.Sprintf("%05d|%x|%X|%o|%b", 42, 255, 255, 8, 5))
	fmt.Println(fmt.Sprintf("|%8.3f|%-8s|%8s|", 3.14159, "left", "right"))
	fmt.Printf("%v %d%%\n", []int{1, 2}, 50)
	// Output:
	// 00042|ff|FF|10|101
	// |   3.142|left    |   right|
	// [1 2] 50%
}
