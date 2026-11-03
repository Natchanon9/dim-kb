// Package format is internal: only packages under packages-modules/ may import it.
package format

import "strings"

// Title upper-cases the first letter.
func Title(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
