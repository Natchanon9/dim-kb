package jsonencoding

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Point struct{ X, Y int }

func Example_gob() {
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(Point{1, 2}) // binary, Go-only format
	var p Point
	gob.NewDecoder(&buf).Decode(&p)
	fmt.Println(p)
	// Output: {1 2}
}
