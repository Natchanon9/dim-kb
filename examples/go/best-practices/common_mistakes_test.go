package bestpractices

import (
	"errors"
	"fmt"
)

type notFound struct{}

func (*notFound) Error() string { return "not found" }

// find returns a literal nil on success, never a typed nil pointer.
func find(ok bool) error {
	if !ok {
		return &notFound{}
	}
	return nil
}

func Example_commonMistakes() {
	m := map[string]int{} // not `var m map[string]int`: writing to a nil map panics
	m["a"]++

	err := find(true)
	var nf *notFound
	fmt.Println(m["a"], err == nil, errors.As(find(false), &nf))
	// Output: 1 true true
}
