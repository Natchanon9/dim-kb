package errorhandling

import (
	"errors"
	"fmt"
)

func Example_errorInterface() {
	// built in: type error interface { Error() string }
	var err error = errors.New("disk full")
	fmt.Println(err, err.Error() == "disk full")
	// Output: disk full true
}
