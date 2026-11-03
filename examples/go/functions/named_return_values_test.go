package functions

import (
	"fmt"
	"strconv"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return: returns x, y
}

func parse(s string) (n int, err error) {
	defer func() { // a deferred func can change named results
		if err != nil {
			err = fmt.Errorf("parse %q: %w", s, err)
		}
	}()
	return strconv.Atoi(s)
}

func Example_namedReturns() {
	fmt.Println(split(17))
	_, err := parse("x")
	fmt.Println(err)
	// Output:
	// 7 10
	// parse "x": strconv.Atoi: parsing "x": invalid syntax
}
