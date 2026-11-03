package errorhandling

import (
	"errors"
	"fmt"
)

func Example_creatingErrors() {
	e1 := errors.New("not found")            // fixed message
	e2 := fmt.Errorf("user %d not found", 7) // formatted message
	fmt.Println(e1)
	fmt.Println(e2)
	// Output:
	// not found
	// user 7 not found
}
