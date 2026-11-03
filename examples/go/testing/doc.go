// Package gotesting holds the snippets for the "Testing" pages
// in src/content/docs/programming/go/testing/.
package gotesting

import (
	"net/http"
	"slices"
)

func Add(a, b int) int { return a + b }

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Reverse(s string) string {
	r := []rune(s)
	slices.Reverse(r)
	return string(r)
}

func Greet(name string) string { return "Hello, " + name + "!" }

func Classify(n int) string {
	switch {
	case n < 0:
		return "negative"
	case n == 0:
		return "zero"
	default:
		return "positive"
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}
