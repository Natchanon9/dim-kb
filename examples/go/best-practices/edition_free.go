//go:build !premium

package bestpractices

// Edition reports which build this is: the default build.
func Edition() string { return "free" }
