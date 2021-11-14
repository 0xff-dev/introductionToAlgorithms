package leetcode

import "testing"

func TestRotateString(t *testing.T) {
	s, g := "abcde", "cdeab"
	if !rotateString(s, g) {
		t.Fatalf("expect true get false")
	}

	s, g = "abcde", "abced"
	if rotateString(s, g) {
		t.Fatalf("expect false get true")
	}

	s, g = "gcmbf", "fgcmb"
	if !rotateString(s, g) {
		t.Fatalf("expect true get false")
	}
}
