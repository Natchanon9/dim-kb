package errorhandling

import (
	"errors"
	"fmt"
)

func Example_bestPractices() {
	err := errors.Join( // Go 1.20+: report several problems at once
		errors.New("name required"),
		errors.New("email invalid"),
	)
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrNotFound))
	// Output:
	// name required
	// email invalid
	// false
}
