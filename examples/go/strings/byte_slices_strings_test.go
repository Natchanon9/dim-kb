package gostrings

import (
	"bytes"
	"fmt"
)

func Example_byteSlices() {
	b := []byte("hello") // copies the string's bytes
	b[0] = 'H'           // []byte is mutable
	fmt.Println(string(b), b, string(bytes.ToUpper(b)))
	// Output: Hello [72 101 108 108 111] HELLO
}
