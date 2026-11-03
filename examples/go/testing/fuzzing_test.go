package gotesting

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	f.Add("hello") // seed corpus: runs on every plain `go test`
	f.Add("สวัสดี")
	f.Fuzz(func(t *testing.T, s string) {
		if !utf8.ValidString(s) {
			t.Skip()
		}
		if got := Reverse(Reverse(s)); got != s {
			t.Errorf("Reverse twice = %q, want %q", got, s)
		}
	})
}
