package fileio

import (
	"bufio"
	"fmt"
	"strings"
)

func Example_bufio() {
	sc := bufio.NewScanner(strings.NewReader("a\nb\nc")) // reads line by line
	n, last := 0, ""
	for sc.Scan() {
		n++
		last = sc.Text()
	}
	fmt.Println(n, last, sc.Err())

	var sb strings.Builder
	w := bufio.NewWriter(&sb)
	w.WriteString("buffered")
	fmt.Println(sb.Len()) // nothing written yet
	w.Flush()
	fmt.Println(sb.String())
	// Output:
	// 3 c <nil>
	// 0
	// buffered
}
