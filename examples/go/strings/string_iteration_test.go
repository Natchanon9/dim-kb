package gostrings

import "fmt"

func Example_iteration() {
	s := "aก"
	for i, r := range s { // by rune
		fmt.Println(i, string(r), r)
	}
	for i := 0; i < len(s); i++ { // by byte
		fmt.Print(s[i], " ")
	}
	// Output:
	// 0 a 97
	// 1 ก 3585
	// 97 224 184 129
}
