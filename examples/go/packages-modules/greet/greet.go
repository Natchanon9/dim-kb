// Package greet builds greeting messages.
package greet

import "github.com/Natchanon9/dim-kb/examples/go/packages-modules/internal/format"

// Hello is exported: other packages can call greet.Hello.
func Hello(name string) string { return format.Title("hello, " + name) }

// suffix is unexported: only code in package greet can use it.
func suffix() string { return "!" }

// Excited uses the unexported helper.
func Excited(name string) string { return Hello(name) + suffix() }
