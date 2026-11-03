package errorhandling

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Field, Reason string
}

func (e *ValidationError) Error() string { return e.Field + ": " + e.Reason }

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return &ValidationError{Field: "email", Reason: "missing @"}
	}
	return nil // a literal nil, never a nil *ValidationError
}

func Example_customErrors() {
	fmt.Println(validateEmail("bob"), validateEmail("a@b.com") == nil)
	// Output: email: missing @ true
}
