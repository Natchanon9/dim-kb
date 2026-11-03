package errorhandling

import (
	"errors"
	"fmt"
)

func Example_wrapping() {
	base := errors.New("connection refused")
	wrapped := fmt.Errorf("query users: %w", base) // %w keeps base inside
	flat := fmt.Errorf("query users: %v", base)    // %v only copies the text
	fmt.Println(wrapped)
	fmt.Println(errors.Unwrap(wrapped) == base, errors.Unwrap(flat) == nil)
	// Output:
	// query users: connection refused
	// true true
}
