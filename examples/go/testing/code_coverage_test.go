package gotesting

import "testing"

func TestClassify(t *testing.T) {
	for in, want := range map[int]string{-1: "negative", 0: "zero", 5: "positive"} {
		if got := Classify(in); got != want {
			t.Errorf("Classify(%d) = %q, want %q", in, got, want)
		}
	}
}
