package gotesting

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		in, want int
	}{
		{"positive", 3, 3},
		{"negative", -3, 3},
		{"zero", 0, 0},
	}
	for _, tt := range tests {
		if got := Abs(tt.in); got != tt.want {
			t.Errorf("%s: Abs(%d) = %d, want %d", tt.name, tt.in, got, tt.want)
		}
	}
}
