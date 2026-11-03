//go:build premium

package bestpractices

// Edition reports which build this is: `go build -tags premium`.
func Edition() string { return "premium" }
