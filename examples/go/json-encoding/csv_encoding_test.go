package jsonencoding

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func Example_csv() {
	r := csv.NewReader(strings.NewReader("name,age\nann,30\n\"b, c\",40\n"))
	rows, _ := r.ReadAll() // handles quoted fields with commas
	fmt.Println(len(rows), rows[2][0])

	var sb strings.Builder
	w := csv.NewWriter(&sb)
	w.Write([]string{"x", "y,z"})
	w.Flush()
	fmt.Print(sb.String())
	// Output:
	// 3 b, c
	// x,"y,z"
}
