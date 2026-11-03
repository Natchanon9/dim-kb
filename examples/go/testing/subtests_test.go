package gotesting

import "testing"

func TestAbsSubtests(t *testing.T) {
	tests := []struct {
		name     string
		in, want int
	}{
		{"positive", 3, 3},
		{"negative", -3, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // go test -run TestAbsSubtests/negative
			t.Parallel()
			if got := Abs(tt.in); got != tt.want {
				t.Errorf("Abs(%d) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
