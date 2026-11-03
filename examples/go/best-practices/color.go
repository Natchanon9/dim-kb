package bestpractices

// Running `go generate ./...` would write color_string.go with the String method.
// The method below is a hand-written equivalent so the example builds without the tool.
//
//go:generate stringer -type=Color

// Color is a small enum.
type Color int

const (
	Red Color = iota
	Green
	Blue
)

// String is what stringer would generate (simplified).
func (c Color) String() string {
	switch c {
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	}
	return "Color(" + itoa(int(c)) + ")"
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	if neg {
		b = append([]byte{'-'}, b...)
	}
	return string(b)
}
