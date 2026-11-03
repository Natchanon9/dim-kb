package gostrings

import (
	"fmt"
	"regexp"
)

var emailRe = regexp.MustCompile(`(\w+)@(\w+)\.com`) // compile once

func Example_regexp() {
	fmt.Println(emailRe.MatchString("a@b.com"))
	fmt.Println(emailRe.FindStringSubmatch("mail bob@example.com now"))
	fmt.Println(emailRe.ReplaceAllString("x@y.com", "$1 at $2"))
	// Output:
	// true
	// [bob@example.com bob example]
	// x at y
}
