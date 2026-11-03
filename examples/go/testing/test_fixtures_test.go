package gotesting

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGreetGolden(t *testing.T) {
	want, err := os.ReadFile(filepath.Join("testdata", "greeting.golden")) // go tooling ignores testdata/
	if err != nil {
		t.Fatal(err)
	}
	if got := Greet("Ann") + "\n"; got != string(want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
