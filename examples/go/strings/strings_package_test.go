package gostrings

import (
	"fmt"
	"strings"
)

func Example_stringsPackage() {
	t := strings.TrimSpace("  Hello, Go  ")
	fmt.Println(strings.ToUpper(t), strings.Contains(t, "Go"), strings.HasPrefix(t, "He"), strings.Index(t, "Go"))
	fmt.Println(strings.Split("a,b,c", ","), strings.Join([]string{"x", "y"}, "-"), strings.ReplaceAll("aaa", "a", "b"), strings.Fields(" a  b "))
	key, value, found := strings.Cut("key=value", "=")
	fmt.Println(strings.EqualFold("Go", "GO"), key, value, found)
	// Output:
	// HELLO, GO true true 7
	// [a b c] x-y bbb [a b]
	// true key value true
}
