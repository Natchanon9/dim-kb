package gostrings

import (
	"fmt"
	"unicode/utf8"
)

func Example_runes() {
	s := "สวัสดี"
	r := []rune(s)
	fmt.Println(len(s), utf8.RuneCountInString(s), r[0] == 'ส', string(r[:2]))
	// Output: 18 6 true สว
}
