package errorhandling

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func Example_isAs() {
	err := fmt.Errorf("get user: %w", ErrNotFound)
	fmt.Println(errors.Is(err, ErrNotFound)) // compares values through the chain

	err = fmt.Errorf("save: %w", validateEmail("bob"))
	var ve *ValidationError
	if errors.As(err, &ve) { // finds a *ValidationError in the chain
		fmt.Println(ve.Field)
	}
	// Output:
	// true
	// email
}
