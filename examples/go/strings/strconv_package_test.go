package gostrings

import (
	"fmt"
	"strconv"
)

func Example_strconv() {
	n, _ := strconv.Atoi("42")
	f, _ := strconv.ParseFloat("3.5", 64)
	b, _ := strconv.ParseBool("true")
	fmt.Println(n+1, f*2, b, strconv.Itoa(7)+"!", strconv.FormatFloat(1.5, 'f', 2, 64), strconv.Quote("hi"))
	_, err := strconv.Atoi("12a")
	fmt.Println(err)
	// Output:
	// 43 7 true 7! 1.50 "hi"
	// strconv.Atoi: parsing "12a": invalid syntax
}
