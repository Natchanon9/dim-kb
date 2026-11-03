package basicsyntax

import "fmt"

type Status int

const (
	Pending Status = iota // 0
	Active                // 1
	Closed                // 2
)

const (
	_  = iota             // skip 0
	KB = 1 << (10 * iota) // 1 << 10
	MB                    // 1 << 20
)

func Example_iota() {
	fmt.Println(Pending, Active, Closed, KB, MB)
	// Output: 0 1 2 1024 1048576
}
