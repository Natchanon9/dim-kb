package gotesting

import (
	"os"
	"testing"
)

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper() // failures point at the caller's line, not this one
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestWithHelpers(t *testing.T) {
	t.Setenv("APP_ENV", "test") // restored after the test
	dir := t.TempDir()
	t.Cleanup(func() { /* close DBs, stop servers ... */ })

	assertEqual(t, Greet("Ann"), "Hello, Ann!")
	assertEqual(t, os.Getenv("APP_ENV"), "test")
	assertEqual(t, dir != "", true)
}
