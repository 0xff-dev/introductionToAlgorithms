package leetcode

import "testing"

func TestHalvesAreAlike(t *testing.T) {
	s := "book"
	if !halvesAreAlike(s) {
		t.Fatalf("expect true get false")
	}

	s = "textbook"
	if halvesAreAlike(s) {
		t.Fatalf("expect false get true")
	}
}
