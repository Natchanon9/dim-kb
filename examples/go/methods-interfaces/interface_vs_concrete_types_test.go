package methodsinterfaces

import "fmt"

type MyErr struct{}

func (*MyErr) Error() string { return "boom" }

func mayFail(fail bool) error {
	var p *MyErr // concrete pointer, nil
	if fail {
		p = &MyErr{}
	}
	return p // BUG: a non-nil interface holding a nil *MyErr
}

func Example_typedNil() {
	err := mayFail(false)
	fmt.Println(err == nil)
	// Output: false
}
