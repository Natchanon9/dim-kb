package gostrings

import "fmt"

func Example_stringBasics() {
	s := "go"
	t := s + "lang"                         // creates a new string
	raw := `C:\path\n`                      // raw string: no escapes
	fmt.Println(t, len("héllo"), s[0], raw) // len counts bytes; s[0] is a byte
	// Output: golang 6 103 C:\path\n
}
